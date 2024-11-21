package service

import (
	"testing"

	"hub-service/model"
	"hub-service/test"

	"github.com/stretchr/testify/assert"
)

func TestAuthenticateByUsernameAndPassword_Success(t *testing.T) {
	container := test.PrepareForServiceTest()

	service := NewUserService(container)
	result, User := service.AuthenticateByUsernameAndPassword("user1", "pw1")

	a := &model.User{}
	a.FindByName(container.GetRepository(), "user1")
	assert.Equal(t, a, User)
	assert.True(t, result)
}

func TestAuthenticateByUsernameAndPassword_EntityNotFound(t *testing.T) {
	container := test.PrepareForServiceTest()

	service := NewUserService(container)
	result, User := service.AuthenticateByUsernameAndPassword("abcde", "abcde")

	assert.Nil(t, User)
	assert.False(t, result)
}

func TestAuthenticateByUsernameAndPassword_AuthenticationFailure(t *testing.T) {
	container := test.PrepareForServiceTest()

	service := NewUserService(container)
	result, User := service.AuthenticateByUsernameAndPassword("user1", "abcde")

	assert.Nil(t, User)
	assert.False(t, result)
}
