package controllers

import "gorm.io/gorm"

func NewDBController(db *gorm.DB) *DBController {
	return &DBController{db}
}

// create database controller
type DBController struct {
	Database *gorm.DB
}
