package model

import (
	"hub-service/config"
	"hub-service/repository"

	"golang.org/x/crypto/bcrypt"
)

// User defines struct of User data.
type User struct {
	Base
	Name     string `json:"name"`
	Password string `json:"-"`
	RoleID   uint64 `json:"role_id"`
	TeamID   uint64 `json:"team_id"`
	Team     *Team  `json:"team,omitempty" gorm:"foreignKey:TeamID"`
	Role     *Role  `json:"role,omitempty" gorm:"foreignKey:RoleID"`
}

// RecordUser defines struct represents the record of the database.
type RecordUser struct {
	ID       uint64
	Name     string
	Password string
	RoleID   uint64
	RoleName string
}

// TableName returns the table name of User struct and it is used by gorm.
func (User) TableName() string {
	return "users"
}

// NewUser is constructor.
func NewUser(name string, password string, roleID uint64) *User {
	return &User{Name: name, Password: password, RoleID: roleID}
}

// NewUserWithPlainPassword is constructor. And it is encoded plain text password by using bcrypt.
func NewUserWithPlainPassword(name string, password string, roleID, teamID uint64) *User {
	hashed, _ := bcrypt.GenerateFromPassword([]byte(password), config.PasswordHashCost)
	return &User{Name: name, Password: string(hashed), RoleID: roleID, TeamID: teamID}
}

// FindByName returns Users full matched given user name.
func (a *User) FindByName(rep repository.Repository, name string) error {
	err := rep.Model(a).First(a, "name = ?", name).Error
	if err != nil {
		return err
	}

	return nil
}

func (a *User) GetUserInfo(rep repository.Repository, name string) error {
	err := rep.Model(a).
		Preload("Role").
		Preload("Team.Hub").
		Find(a, "name = ?", name).
		Error
	if err != nil {
		return err
	}

	return nil
}

// Create persists this User data.
func (a *User) Create(rep repository.Repository) (*User, error) {
	if err := rep.Create(a).Error; err != nil {
		return nil, err
	}
	return a, nil
}

// ToString is return string of object
func (a *User) ToString() string {
	return toString(a)
}
