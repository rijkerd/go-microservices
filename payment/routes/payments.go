package routes

import (
	"payment/controllers"

	"github.com/gofiber/fiber/v2"
)

func PaymentRoute(route fiber.Router) {
	route.Get("/", controllers.GetAllPayments)
	route.Get("/:id", controllers.GetSinglePayment)
}
