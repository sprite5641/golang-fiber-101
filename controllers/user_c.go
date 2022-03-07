package controllers

import (
	"errors"
	config "go-fiber-app/configs"
	"go-fiber-app/models"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func (db *DBController) Login(c *fiber.Ctx) error {

	var input models.Login

	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Error on login request", "data": err})
	}
	identity := input.PhoneNumber
	pass := input.Password

	user, err := getUserByPhoneNumber(identity, db)
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

func getUserByPhoneNumber(p string, db *DBController) (*models.User, error) {
	var user models.User
	if err := db.Database.Where(&models.User{PhoneNumber: p}).Find(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (db *DBController) GetUsers(c *fiber.Ctx) error {

	var user []models.User
	// find all user in the database
	db.Database.Find(&user)

	// If no note is present return an error
	if len(user) == 0 {
		return c.Status(404).JSON(fiber.Map{"success": false, "message": "No user present", "data": nil})
	}

	// Else return user
	return c.JSON(fiber.Map{"success": true, "message": "user Found", "data": user})
}

func (db *DBController) Register(c *fiber.Ctx) error {

	user := new(models.Users)

	// Store the body in the book and return error if encountered
	err := c.BodyParser(user)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"success": false, "message": "Review your input", "data": err})
	}

	hashed, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 10)

	user.Password = string(hashed)

	// Create the book and return error if encountered
	err = db.Database.Create(&user).Error
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

func (db *DBController) GetUserById(c *fiber.Ctx) error {
	var user models.User

	user_id := c.Params("user_id")

	db.Database.First(&user, user_id)

	return c.JSON(fiber.Map{"success": true, "message": "Created user", "data": user})
}
