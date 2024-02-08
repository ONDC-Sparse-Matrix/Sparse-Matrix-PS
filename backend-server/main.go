package main

import (
    "backend-server/configs"
    "backend-server/routes" //add this
    "github.com/gofiber/fiber/v2"
)

func main() {
    app := fiber.New()

    configs.ConnectDB()

    routes.MerchantRoute(app)

    app.Listen(":4000")
}