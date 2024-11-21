package controller

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"hub-service/config"
	"hub-service/model"
	"hub-service/model/dto"
	"hub-service/test"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestGetLoginStatus_Success(t *testing.T) {
	router, container := test.PrepareForControllerTest(false)

	user := NewUserController(container)
	router.GET(config.APIUserLoginStatus, func(c echo.Context) error { return user.GetLoginStatus(c) })

	req := httptest.NewRequest("GET", config.APIUserLoginStatus, nil)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.JSONEq(t, "true", rec.Body.String())
}

func TestGetLoginUser_Success(t *testing.T) {
	router, container := test.PrepareForControllerTest(false)

	user := NewUserController(container)
	router.GET(config.APIUserLoginUser, func(c echo.Context) error { return user.GetLoginUser(c) })

	req := httptest.NewRequest("GET", config.APIUserLoginUser, nil)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)
	assert.Equal(t, http.StatusOK, rec.Code)
}

func TestLogin_Success(t *testing.T) {
	router, container := test.PrepareForControllerTest(true)

	user := NewUserController(container)
	router.POST(config.APIUserLogin, func(c echo.Context) error { return user.Login(c) })

	param := createLoginSuccessUser()
	req := test.NewJSONRequest("POST", config.APIUserLogin, param)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.NotEmpty(t, test.GetCookie(rec, "GSESSION"))
}

func TestLogin_AuthenticationFailure(t *testing.T) {
	router, container := test.PrepareForControllerTest(true)

	user := NewUserController(container)
	router.POST(config.APIUserLogin, func(c echo.Context) error { return user.Login(c) })

	param := createLoginFailureUser()
	req := test.NewJSONRequest("POST", config.APIUserLogin, param)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusUnauthorized, rec.Code)
	assert.Empty(t, test.GetCookie(rec, "GSESSION"))
}

func TestLogout_Success(t *testing.T) {
	router, container := test.PrepareForControllerTest(true)

	user := NewUserController(container)
	router.POST(config.APIUserLogout, func(c echo.Context) error { return user.Logout(c) })

	req := test.NewJSONRequest("POST", config.APIUserLogout, nil)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.NotEmpty(t, test.GetCookie(rec, "GSESSION"))
}

func TestGetUserInfo_Success(t *testing.T) {
	router, container := test.PrepareForControllerTest(true)

	user := NewUserController(container)

	a := &model.User{}
	a.FindByName(container.GetRepository(), "user1")
	router.GET(config.APIGetUserInfo, func(c echo.Context) error {
		sess := container.GetSession()
		_ = sess.SetUser(c, a)
		_ = sess.Save(c)
		return user.GetUserInfo(c)
	})

	req := test.NewJSONRequest("GET", config.APIGetUserInfo+"?name=user1", nil)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	var userReq model.User
	json.Unmarshal(rec.Body.Bytes(), &userReq)

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, userReq.Name, a.Name)
	assert.Equal(t, userReq.Team.Name, "teamtest")
	assert.Equal(t, userReq.Team.Hub.Name, "hubtest")

}

func createLoginSuccessUser() *dto.LoginDto {
	return &dto.LoginDto{
		UserName: "user1",
		Password: "pw1",
	}
}

func createLoginFailureUser() *dto.LoginDto {
	return &dto.LoginDto{
		UserName: "user1",
		Password: "abcde",
	}
}
