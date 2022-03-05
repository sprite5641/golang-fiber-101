package routes

import (
	"go-fiber-app/controllers"
	"go-fiber-app/middleware"

	"github.com/gofiber/fiber/v2"
)

func UserRoute(route fiber.Router) {

	route.Get("", middleware.Protected(), controllers.GetUsers)
	route.Post("", controllers.Register)
	route.Post("/login", controllers.Login)

}
