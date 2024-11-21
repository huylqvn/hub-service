package router

import (
	"net/http"

	"hub-service/config"
	"hub-service/container"
	"hub-service/controller"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	_ "hub-service/docs" // for using echo-swagger

	echoSwagger "github.com/swaggo/echo-swagger"
)

// Init initialize the routing of this application.
func Init(e *echo.Echo, container container.Container) {
	setCORSConfig(e, container)

	setErrorController(e, container)
	setUserController(e, container)
	setHealthController(e, container)

	setSwagger(container, e)
}

func setCORSConfig(e *echo.Echo, container container.Container) {
	if container.GetConfig().Extension.CorsEnabled {
		e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
			AllowCredentials:                         true,
			UnsafeWildcardOriginWithAllowCredentials: true,
			AllowOrigins:                             []string{"*"},
			AllowHeaders: []string{
				echo.HeaderAccessControlAllowHeaders,
				echo.HeaderContentType,
				echo.HeaderContentLength,
				echo.HeaderAcceptEncoding,
			},
			AllowMethods: []string{
				http.MethodGet,
				http.MethodPost,
				http.MethodPut,
				http.MethodDelete,
			},
			MaxAge: 86400,
		}))
	}
}

func setErrorController(e *echo.Echo, container container.Container) {
	errorHandler := controller.NewErrorController(container)
	e.HTTPErrorHandler = errorHandler.JSONError
	e.Use(middleware.Recover())
}

func setUserController(e *echo.Echo, container container.Container) {
	user := controller.NewUserController(container)
	e.GET(config.APIUserLoginStatus, func(c echo.Context) error { return user.GetLoginStatus(c) })
	e.GET(config.APIUserLoginUser, func(c echo.Context) error { return user.GetLoginUser(c) })

	if container.GetConfig().Extension.SecurityEnabled {
		e.POST(config.APIUserLogin, func(c echo.Context) error { return user.Login(c) })
		e.POST(config.APIUserLogout, func(c echo.Context) error { return user.Logout(c) })
		e.GET(config.APIGetUserInfo, func(c echo.Context) error { return user.GetUserInfo(c) })
	}
}

func setHealthController(e *echo.Echo, container container.Container) {
	health := controller.NewHealthController(container)
	e.GET(config.APIHealth, func(c echo.Context) error { return health.GetHealthCheck(c) })
}

func setSwagger(container container.Container, e *echo.Echo) {
	if container.GetConfig().Swagger.Enabled {
		e.GET("/swagger/*", echoSwagger.WrapHandler)
	}
}
