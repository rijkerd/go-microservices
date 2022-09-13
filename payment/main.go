package main

import (
	"log"
	"os"
	"payment/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"

	_ "payment/docs"

	"github.com/gofiber/swagger"

	"github.com/joho/godotenv"
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

// @title Payment
// @version 1.0
// @description Payment api build with go Fiber
// @termsOfService http://rijkerd.github.io
// @contact.name Richard Aggrey
// @contact.email richardaggrey7@gmail.com
// @contact.github https://github.com/rijkerd
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:3000
// @BasePath /
func main() {
	if os.Getenv("APP_ENV") != "production" {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}
	app := fiber.New()

	app.Use(cors.New())
	app.Use(logger.New())

	app.Get("/metrics", monitor.New(monitor.Config{Title: "Payment Service"}))
	app.Get("/swagger/*", swagger.HandlerDefault)

	setupRoutes(app)

	port := os.Getenv("PORT")
	err := app.Listen(":" + port)

	if err != nil {
		log.Fatal("Error app failed to start")
		panic(err)
	}
}
