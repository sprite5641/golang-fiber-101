package routes

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func IndexRoute(route fiber.Router, db *gorm.DB) {
	r := route.Group("/users")
	UserRoute(r, db)
}
