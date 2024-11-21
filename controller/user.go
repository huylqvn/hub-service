package controller

import (
	"net/http"

	"hub-service/container"
	"hub-service/model"
	"hub-service/model/dto"
	"hub-service/service"

	"github.com/labstack/echo/v4"
)

// UserController is a controller for managing user User.
type UserController interface {
	GetLoginStatus(c echo.Context) error
	GetLoginUser(c echo.Context) error
	Login(c echo.Context) error
	Logout(c echo.Context) error
	GetUserInfo(c echo.Context) error
}

type userController struct {
	context container.Container
	service service.UserService
}

// NewUserController is constructor.
func NewUserController(container container.Container) UserController {
	return &userController{
		context: container,
		service: service.NewUserService(container),
	}
}

// GetLoginStatus returns the status of login.
// @Summary Get the login status.
// @Description Get the login status of current logged-in user.
// @Tags Auth
// @Accept  json
// @Produce  json
// @Success 200 {boolean} bool "The current user have already logged-in. Returns true."
// @Failure 401 {boolean} bool "The current user haven't logged-in yet. Returns false."
// @Router /auth/loginStatus [get]
func (controller *userController) GetLoginStatus(c echo.Context) error {
	return c.JSON(http.StatusOK, true)
}

// GetLoginUser returns the User data of logged in user.
// @Summary Get the User data of logged-in user.
// @Description Get the User data of logged-in user.
// @Tags Auth
// @Accept  json
// @Produce  json
// @Success 200 {object} model.User "Success to fetch the User data. If the security function is disable, it returns the dummy data."
// @Failure 401 {boolean} bool "The current user haven't logged-in yet. Returns false."
// @Router /auth/loginUser [get]
func (controller *userController) GetLoginUser(c echo.Context) error {
	return c.JSON(http.StatusOK, controller.context.GetSession().GetUser(c))
}

// Login is the method to login using username and password by http post.
// @Summary Login using username and password.
// @Description Login using username and password.
// @Tags Auth
// @Accept  json
// @Produce  json
// @Param data body dto.LoginDto true "User name and Password for logged-in."
// @Success 200 {object} model.User "Success to the authentication."
// @Failure 401 {boolean} bool "Failed to the authentication."
// @Router /auth/login [post]
func (controller *userController) Login(c echo.Context) error {
	dto := dto.NewLoginDto()
	if err := c.Bind(dto); err != nil {
		return c.JSON(http.StatusBadRequest, dto)
	}

	sess := controller.context.GetSession()
	if user := sess.GetUser(c); user != nil {
		return c.JSON(http.StatusOK, user)
	}

	authenticate, a := controller.service.AuthenticateByUsernameAndPassword(dto.UserName, dto.Password)
	if authenticate {
		_ = sess.SetUser(c, a)
		_ = sess.Save(c)
		return c.JSON(http.StatusOK, a)
	}
	return c.NoContent(http.StatusUnauthorized)
}

// Logout is the method to logout by http post.
// @Summary Logout.
// @Description Logout.
// @Tags Auth
// @Accept  json
// @Produce  json
// @Success 200
// @Router /auth/logout [post]
func (controller *userController) Logout(c echo.Context) error {
	sess := controller.context.GetSession()
	_ = sess.SetUser(c, nil)
	_ = sess.Delete(c)
	return c.NoContent(http.StatusOK)
}

// GetUserInfo returns the User data of logged in user.
// @Summary Get the User data of logged-in user.
// @Description Get the User data of logged-in user.
// @Tags Auth
// @Accept  json
// @Produce  json
// @Success 200 {object} model.User "Success to fetch the User data."
// @Failure 401 {boolean} bool "The current user haven't logged-in yet. Returns false."
// @Router /auth/getUserInfo [get]
func (controller *userController) GetUserInfo(c echo.Context) error {
	sess := controller.context.GetSession()
	if user := sess.GetUser(c); user == nil {
		return c.JSON(http.StatusUnauthorized, dto.AuthenticationFail{
			Message: "Unauthenticated",
		})
	}

	dto := dto.GetUserInfo{}
	if err := c.Bind(&dto); err != nil {
		return c.JSON(http.StatusBadRequest, dto)
	}

	user := &model.User{}
	err := user.GetUserInfo(controller.context.GetRepository(), dto.Name)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto)
	}
	return c.JSON(http.StatusOK, user)
}
