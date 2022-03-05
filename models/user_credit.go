package models

import (
	"time"

	"gorm.io/gorm"
)

type UserCredits struct {
	gorm.Model
	ID uint `gorm:"primary key;autoIncrement" json:"id"`
	// UserId    Users
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}
