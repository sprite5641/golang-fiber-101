package controllers

import (
	"errors"
	config "go-fiber-app/configs"
	"go-fiber-app/database"
	"go-fiber-app/models"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func getUserByPhoneNumber(p string) (*models.Users, error) {
	db := database.DB
	var user models.Users
	if err := db.Where(&models.Users{PhoneNumber: p}).Find(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func Login(c *fiber.Ctx) error {
	type LoginInput struct {
		PhoneNumber string `json:"phone_number"`
		Password    string `json:"password"`
	}

	var input LoginInput

	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Error on login request", "data": err})
	}
	identity := input.PhoneNumber
	pass := input.Password

	user, err := getUserByPhoneNumber(identity)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "error", "message": "Error on username", "data": err})
	}

	if user == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "error", "message": "User not found", "data": err})
	}

	if !CheckPasswordHash(pass, user.Password) {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "error", "message": "Invalid password", "data": nil})
	}

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["phone_number"] = user.PhoneNumber
	claims["user_id"] = user.ID
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	t, err := token.SignedString([]byte(config.Config("SECRET")))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Success login", "data": t})
}

func GetUsers(c *fiber.Ctx) error {
	db := database.DB
	var user []models.User

	// find all user in the database
	db.Find(&user)

	// If no note is present return an error
	if len(user) == 0 {
		return c.Status(404).JSON(fiber.Map{"success": false, "message": "No user present", "data": nil})
	}

	// Else return user
	return c.JSON(fiber.Map{"success": true, "message": "user Found", "data": user})
}

func Register(c *fiber.Ctx) error {
	db := database.DB
	user := new(models.Users)

	// Store the body in the book and return error if encountered
	err := c.BodyParser(user)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"success": false, "message": "Review your input", "data": err})
	}

	hashed, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 10)

	user.Password = string(hashed)

	// Create the book and return error if encountered
	err = db.Create(&user).Error
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"success": false, "message": "Could not create user", "data": err})
	}

	// Return the created user
	return c.JSON(fiber.Map{"success": true, "message": "Created user", "data": user})
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
