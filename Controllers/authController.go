package controllers

import (
	"encoding/json"
	"os"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"github.com/vitaly06/sales-rest-api/Models"
	db "github.com/vitaly06/sales-rest-api/config"
)

func Login(c *fiber.Ctx) error {
	cashierId, err := c.ParamsInt("cashierId")

	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": "Cashier id must be a number",
		})
	}

	var data map[string]interface{}

	if err := json.Unmarshal(c.Body(), &data); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": "Invalid data",
			"error":   err.Error(),
		})
	}

	password, ok := data["password"].(string)
	if !ok || password == "" {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": "Password is required",
		})
	}

	var cashier Models.Cashier

	db.DB.Where("id=?", cashierId).First(&cashier)

	if cashier.Id == 0 {
		return c.Status(404).JSON(fiber.Map{
			"success": false,
			"message": "Cashier not found",
		})
	}

	if cashier.Password != password {
		return c.Status(401).JSON(fiber.Map{
			"success": false,
			"message": "Passwords don't match",
		})
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"Issuer":    strconv.Itoa(int(cashier.Id)),
		"ExpiresAt": time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"success": false,
			"message": "Could not generate token",
		})
	}

	cashierData := make(map[string]interface{})

	cashierData["token"] = tokenString

	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"message": "Successfully authorization",
		"data":    cashierData,
	})
}

func Logout(c *fiber.Ctx) error {
	return nil
}

func Password(c *fiber.Ctx) error {
	return nil
}
