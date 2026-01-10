package routes

import (
	Controllers "github.com/vitaly06/sales-rest-api/Controllers"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	// app.Post("/cashiers/:cashierId/login", controller.Login)
	// app.Get("/cashiers/:cashierId/logout", controller.Logout)
	// app.Post("/cashiers/:cashierId/password", controller.Password)

	// cashier routes
	app.Post("/cashiers", Controllers.CreateCashier)
	app.Get("/cashiers", Controllers.CashiersList)
	app.Get("/cashiers/:cashierId", Controllers.GetCashierDetails)
	app.Delete("/cashiers/:cashierId", Controllers.DeleteCashier)
	app.Put("/cashiers/:cashierId", Controllers.UpdateCashier)
}
