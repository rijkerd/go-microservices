package controllers

import (
	"payment/managers"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func GetAllPayments(c *fiber.Ctx) error {
	payments := managers.GetPayments()
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": payments,
	})
}

func GetSinglePayment(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	payment := managers.GetPaymentByID(id)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": payment,
	})
}
