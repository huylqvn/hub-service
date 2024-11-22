package controller

import (
	"net/http"
	"net/http/httptest"
	"testing"

	_ "hub-service/docs" // for using echo-swagger
	"hub-service/test"

	"github.com/stretchr/testify/assert"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func TestSwagger(t *testing.T) {
	router, _ := test.PrepareForControllerTest(false)
	router.GET("/swagger/*", echoSwagger.WrapHandler)

	req := httptest.NewRequest("GET", "/swagger/index.html", nil)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Regexp(t, "Swagger UI", rec.Body.String())
}
