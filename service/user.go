package service

import (
	"hub-service/container"
	"hub-service/model"

	"golang.org/x/crypto/bcrypt"
)

// UserService is a service for managing user User.
type UserService interface {
	AuthenticateByUsernameAndPassword(username string, password string) (bool, *model.User)
}

type userService struct {
	container container.Container
}

// NewUserService is constructor.
func NewUserService(container container.Container) UserService {
	return &userService{container: container}
}

// AuthenticateByUsernameAndPassword authenticates by using username and plain text password.
func (a *userService) AuthenticateByUsernameAndPassword(username string, password string) (bool, *model.User) {
	rep := a.container.GetRepository()
	logger := a.container.GetLogger()
	user := &model.User{}
	err := user.FindByName(rep, username)
	if err != nil {
		logger.GetZapLogger().Info("User not found: " + err.Error())
		return false, nil
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		logger.GetZapLogger().Info("authentication failure: " + err.Error())
		return false, nil
	}

	return true, user
}
