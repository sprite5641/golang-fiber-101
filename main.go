package main

import (
	database "go-fiber-app/database"
	"go-fiber-app/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func setupRoutes(app *fiber.App) {

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

	routes.IndexRoute(api.Group("/v1"))

}

func main() {
	app := fiber.New()
	app.Use(logger.New())
	database.ConnectDB()
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "https://gofiber.io, https://gofiber.net",
		AllowHeaders:     "Origin, Content-Type, Option, Authorization, X-Session-Id",
		AllowCredentials: true,
	}))

	setupRoutes(app)

	err := app.Listen("127.0.0.1:8080")

	if err != nil {
		panic(err)
	}
}
