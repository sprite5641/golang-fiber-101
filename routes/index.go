package routes

import (
	"github.com/gofiber/fiber/v2"
)

func IndexRoute(route fiber.Router) {
	UserRoute(route.Group("/users"))
}
