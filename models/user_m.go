package models

import (
	"time"

	"gorm.io/gorm"
)

// user migrate
type Users struct {
	ID          uint   `gorm:"primaryKey"`
	PhoneNumber string `gorm:"unique,TYPE:STRING(10)" json:"phone_number"`
	Password    string `json:"password"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

type Login struct {
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
}

type User struct {
	ID          uint   `json:"id"`
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
}
