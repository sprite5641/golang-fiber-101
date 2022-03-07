package main

import (
	"fmt"
	database "go-fiber-app/database"
	"go-fiber-app/routes"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"gorm.io/gorm"
)

func setupRoutes(app *fiber.App, db *gorm.DB) {

	// moved from main method
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"success": true,
			"message": "You are at the endpoint 😉",
		})
	})

	// api group
	api := app.Group("/api")

	// give response when at /api
	api.Get("", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"success": true,
			"message": "You are at the api endpoint 😉",
		})
	})

	v1 := api.Group("/v1")

	routes.IndexRoute(v1, db)

}

func main() {
	app := fiber.New()
	app.Use(logger.New())
	DB := database.ConnectDB()
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:3000",
		AllowHeaders:     "Origin, Content-Type, Option, Authorization, X-Session-Id",
		AllowCredentials: true,
	}))

	setupRoutes(app, DB)

	PORT := os.Getenv("PORT")
	port := fmt.Sprintf("127.0.0.1:%v", PORT)

	fmt.Println("Server Running on Port", port)

	err := app.Listen(port)

	if err != nil {
		panic(err)
	}
}
