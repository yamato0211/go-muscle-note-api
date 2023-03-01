package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (base *Base) BeforeCreate(tx *gorm.DB) (err error) {
	base.ID = uuid.New().String()
	return
}

type Base struct {
	ID        string    `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type User struct {
	Base
	Name         string `gorm:"not null" json:"name"`
	Email        string `gorm:"unique;not null" json:"email"`
	PasswordHash string `gorm:"not null" json:"-"`
}
