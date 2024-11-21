package middleware

import (
	"io"
	"net/http"
	"regexp"
	"strconv"

	"hub-service/container"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/valyala/fasttemplate"
)

// InitLoggerMiddleware initialize a middleware for logger.
func InitLoggerMiddleware(e *echo.Echo, container container.Container) {
	e.Use(RequestLoggerMiddleware(container))
	e.Use(ActionLoggerMiddleware(container))
}

// InitSessionMiddleware initialize a middleware for session management.
func InitSessionMiddleware(e *echo.Echo, container container.Container) {
	conf := container.GetConfig()

	e.Use(session.Middleware(container.GetSession().GetStore()))
	if conf.Extension.SecurityEnabled {
		e.Use(AuthenticationMiddleware(container))
	}
}

// RequestLoggerMiddleware is middleware for logging the contents of requests.
func RequestLoggerMiddleware(container container.Container) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			req := c.Request()
			res := c.Response()
			if err := next(c); err != nil {
				c.Error(err)
			}

			template := fasttemplate.New(container.GetConfig().Log.RequestLogFormat, "${", "}")
			logstr := template.ExecuteFuncString(func(w io.Writer, tag string) (int, error) {
				switch tag {
				case "remote_ip":
					return w.Write([]byte(c.RealIP()))
				case "User_name":
					if User := container.GetSession().GetUser(c); User != nil {
						return w.Write([]byte(User.Name))
					}
					return w.Write([]byte("None"))
				case "uri":
					return w.Write([]byte(req.RequestURI))
				case "method":
					return w.Write([]byte(req.Method))
				case "status":
					return w.Write([]byte(strconv.Itoa(res.Status)))
				default:
					return w.Write([]byte(""))
				}
			})
			container.GetLogger().GetZapLogger().Infof(logstr)
			return nil
		}
	}
}

// ActionLoggerMiddleware is middleware for logging the start and end of controller processes.
// ref: https://echo.labstack.com/cookbook/middleware
func ActionLoggerMiddleware(container container.Container) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			logger := container.GetLogger()
			logger.GetZapLogger().Debugf(c.Path() + " Action Start")
			if err := next(c); err != nil {
				c.Error(err)
			}
			logger.GetZapLogger().Debugf(c.Path() + " Action End")
			return nil
		}
	}
}

// AuthenticationMiddleware is the middleware of session authentication for echo.
func AuthenticationMiddleware(container container.Container) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if !hasAuthorization(c, container) {
				return c.JSON(http.StatusUnauthorized, false)
			}
			if err := next(c); err != nil {
				c.Error(err)
			}
			return nil
		}
	}
}

// hasAuthorization judges whether the user has the right to access the path.
func hasAuthorization(c echo.Context, container container.Container) bool {
	currentPath := c.Path()
	if equalPath(currentPath, container.GetConfig().Security.AuthPath) {
		if equalPath(currentPath, container.GetConfig().Security.ExculdePath) {
			return true
		}
		User := container.GetSession().GetUser(c)
		if User == nil {
			return false
		}
		if User.Role.Name == "Admin" && equalPath(currentPath, container.GetConfig().Security.AdminPath) {
			_ = container.GetSession().Save(c)
			return true
		}
		if User.Role.Name == "User" && equalPath(currentPath, container.GetConfig().Security.UserPath) {
			_ = container.GetSession().Save(c)
			return true
		}
		return false
	}
	return true
}

// equalPath judges whether a given path contains in the path list.
func equalPath(cpath string, paths []string) bool {
	for i := range paths {
		if regexp.MustCompile(paths[i]).Match([]byte(cpath)) {
			return true
		}
	}
	return false
}
