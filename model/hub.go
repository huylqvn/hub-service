package model

import "hub-service/repository"

// Hub defines struct of hub data.
type Hub struct {
	Base
	Name  string `json:"name" gorm:"index:hub_idx_name,unique"`
	Teams []Team `json:"teams" gorm:"-"`
}

// TableName returns the table name of Hub struct and it is used by gorm.
func (Hub) TableName() string {
	return "hubs"
}

// NewHub is constructor.
func NewHub(name string) *Hub {
	return &Hub{Name: name}
}

// Create persists this Hub data.
func (a *Hub) Create(rep repository.Repository) (*Hub, error) {
	if err := rep.Create(a).Error; err != nil {
		return nil, err
	}
	return a, nil
}

// find hub by name
func (a *Hub) FindByName(rep repository.Repository, name string) (*Hub, error) {
	var hub *Hub
	db := rep.Model(hub)
	if err := db.Where("name = ?", name).First(&hub).Error; err != nil {
		return nil, err
	}
	return hub, nil
}

// ToString is return string of object
func (a *Hub) ToString() string {
	return toString(a)
}
