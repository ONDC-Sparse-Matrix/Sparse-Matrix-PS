package routes

import (
    "backend-server/controllers"
    "github.com/gofiber/fiber/v2"
)

func MerchantRoute(app *fiber.App) {
    app.Get("/merchant/:pinCode",controllers.GetMerchants)
	app.Post("/merchant",controllers.AddMerchants)
}
