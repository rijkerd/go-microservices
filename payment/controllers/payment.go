package controllers

import (
	"payment/managers"

	"github.com/gofiber/fiber/v2"
)

func GetAllPayments(c *fiber.Ctx) error {
	payments := managers.GetPayments()
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": payments,
	})
}

// TODO: Add Get Payment Based on ID
