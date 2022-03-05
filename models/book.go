package models

import "time"

type Books struct {
	ID        uint       `gorm:"primary key;autoIncrement" json:"id"`
	Author    *string    `json:"author"`
	Title     *string    `json:"title"`
	Publisher *string    `json:"publisher"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}
