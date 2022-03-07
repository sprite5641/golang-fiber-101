package database

import (
	"fmt"
	config "go-fiber-app/configs"
	"go-fiber-app/models"
	"log"
	"os"
	"strconv"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Declare the variable for the database
var DB *gorm.DB

// ConnectDB connect to db
func ConnectDB() *gorm.DB {
	DB_PORT := config.Config("DB_PORT")
	port, err := strconv.ParseUint(DB_PORT, 10, 32)

	if err != nil {
		log.Println("Idiot")
	}

	DB_HOST := os.Getenv("DB_HOST")
	DB_NAME := os.Getenv("DB_NAME")
	DB_USER := os.Getenv("DB_USER")
	DB_PASSWORD := os.Getenv("DB_PASSWORD")

	// Connection URL to connect to Postgres Database
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", DB_HOST, port, DB_USER, DB_PASSWORD, DB_NAME)
	// Connect to the DB and initialize the DB variable

	DB, err = gorm.Open(postgres.Open(dsn))

	if err != nil {
		panic("failed to connect database")
	}

	fmt.Println("Connection Opened to Database")

	DB.AutoMigrate(&models.Books{}, &models.Users{})
	fmt.Println("Database Migrated")

	return DB
}
