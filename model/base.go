package model

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Base struct {
	ID        uint64     `json:"id,omitempty" gorm:"primaryKey"`
	UUID      string     `json:"uuid" gorm:"type:varchar(36);uniqueIndex;not null"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`

	DeletedAt *time.Time `json:"deleted_at"`
}

func (b *Base) BeforeCreate(tx *gorm.DB) (err error) {
	now := time.Now()

	b.UUID = uuid.New().String()

	b.CreatedAt = &now
	b.UpdatedAt = &now
	return nil
}

func (b *Base) BeforeUpdate(tx *gorm.DB) (err error) {
	now := time.Now()
	b.UpdatedAt = &now
	return nil
}

// DomainObject defines the common interface for domain models.
type DomainObject interface {
	User | Role | Hub | Team
}

// toString returns the JSON data of the domain models.
func toString[T DomainObject](o *T) string {
	var bytes []byte
	var err error
	if bytes, err = json.Marshal(o); err != nil {
		return ""
	}
	return string(bytes)
}
