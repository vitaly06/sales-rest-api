package main

import (
	routes "github.com/vitaly06/sales-rest-api/Routes"
	db "github.com/vitaly06/sales-rest-api/config"

	"fmt"

	"github.com/gofiber/fiber/v2"
)

func main() {
	fmt.Println("Started...")

	db.Connect()

	app := fiber.New()

	routes.Setup(app)

	app.Listen(":3000")
}
