package controllers

import (
	"time"

	Models "github.com/vitaly06/sales-rest-api/Models"
	db "github.com/vitaly06/sales-rest-api/config"

	"github.com/gofiber/fiber/v2"
)

func CreateCashier(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": "Invalid data",
		})
	}

	if data["name"] == "" {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": "Cashier Name is required",
		})
	}

	if data["password"] == "" {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": "Cashier password is required",
		})
	}

	cashier := Models.Cashier{
		Name:      data["name"],
		Password:  data["password"],
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
	}

	db.DB.Create(&cashier)

	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"message": "Cashier added successfully",
		"data":    cashier,
	})
}

func UpdateCashier(c *fiber.Ctx) error {
	return nil
}

func DeleteCashier(c *fiber.Ctx) error {
	return nil
}

func CashiersList(c *fiber.Ctx) error {
	return c.SendString("cashiers list")
}

func GetCashierDetails(c *fiber.Ctx) error {
	return nil
}
