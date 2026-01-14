package controllers

import (
	"fmt"
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
	var cashier []Models.Cashier

	limit := c.QueryInt("limit")
	skip := c.QueryInt("skip")

	var count int64

	db.DB.Select("*").Limit(limit).Offset(skip).Find(&cashier).Count(&count)

	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"message": "Cashier list api",
		"data":    cashier,
	})
}

func GetCashierDetails(c *fiber.Ctx) error {
	cashierId, err := c.ParamsInt("cashierId")
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": "Cashier id must be a number",
		})
	}

	var cashier Models.Cashier

	db.DB.Select("id,name,created_at,updated_at").Where("id=?", cashierId).First(&cashier)

	if cashier.Id == 0 {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": "Cashier not found",
		})
	}

	cashierData := make(map[string]interface{})
	cashierData["id"] = cashier.Id
	cashierData["name"] = cashier.Name
	cashierData["createdAt"] = cashier.CreatedAt
	cashierData["updatedAt"] = cashier.UpdatedAt

	fmt.Println(cashierData)

	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"message": "success",
		"data":    cashierData,
	})
}
