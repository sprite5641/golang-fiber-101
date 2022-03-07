package routes

import (
	"go-fiber-app/controllers"
	"go-fiber-app/middleware"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func UserRoute(route fiber.Router, db *gorm.DB) {

	ctrls := controllers.NewDBController(db)

	route.Get("", middleware.Protected(), ctrls.GetUsers)
	route.Get(":user_id", middleware.Protected(), ctrls.GetUserById)

	route.Post("", ctrls.Register)
	route.Post("/login", ctrls.Login)

}
