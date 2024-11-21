package model

import (
	"hub-service/repository"
)

// Role defines struct of Role data.
type Role struct {
	Base
	Name string `json:"name"`
}

// TableName returns the table name of Role struct and it is used by gorm.
func (Role) TableName() string {
	return "roles"
}

// NewRole is constructor.
func NewRole(name string) *Role {
	return &Role{Name: name}
}

// Create persists this Role data.
func (a *Role) Create(rep repository.Repository) (*Role, error) {
	if err := rep.Create(a).Error; err != nil {
		return nil, err
	}
	return a, nil
}

// ToString is return string of object
func (a *Role) ToString() string {
	return toString(a)
}
