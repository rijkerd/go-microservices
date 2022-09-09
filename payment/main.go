package main

import (
	"payment/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func setupRoutes(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"success":     true,
			"message":     "You are at the root endpoint 😉",
			"github_repo": "rijkerd",
		})
	})

	api := app.Group("/api")

	routes.PaymentRoute(api.Group("/payment"))
}

func main() {
	app := fiber.New()

	app.Use(cors.New())
	app.Use(logger.New())

	setupRoutes(app)

	app.Listen(":3000")
}
