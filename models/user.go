package models

import (
	"time"

	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	ID          uint           `gorm:"primary key;autoIncrement" json:"id"`
	PhoneNumber string         `gorm:"unique,TYPE:STRING(10)" json:"phone_number"`
	Password    string         `json:"password"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at"`
}

type User struct {
	gorm.Model
	ID          uint   `gorm:"primary key;autoIncrement" json:"id"`
	PhoneNumber string `gorm:"unique,TYPE:STRING(10)" json:"phone_number"`
}
