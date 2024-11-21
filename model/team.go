package model

import "hub-service/repository"

type Team struct {
	Base
	Name  string `json:"name" gorm:"index:team_idx_name,unique"`
	HubID uint64 `json:"hub_id" gorm:"index:team_idx_hub_id"`
	Users []User `json:"users" gorm:"-"`
	Hub   *Hub   `json:"hub,omitempty" gorm:"foreignKey:HubID"`
}

func (Team) TableName() string {
	return "teams"
}

func NewTeam(name string, hubID uint64) *Team {
	return &Team{Name: name, HubID: hubID}
}

func (a *Team) Create(rep repository.Repository) (*Team, error) {
	if err := rep.Create(a).Error; err != nil {
		return nil, err
	}
	return a, nil
}

func (a *Team) FindByName(rep repository.Repository, name string) (*Team, error) {
	var team *Team
	db := rep.Model(team)
	if err := db.Where("name = ?", name).First(&team).Error; err != nil {
		return nil, err
	}
	return team, nil
}

func (a *Team) ToString() string {
	return toString(a)
}
