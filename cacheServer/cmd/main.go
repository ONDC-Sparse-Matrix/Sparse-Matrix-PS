package main

import (
	"cache-server/utils"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func Initialize() {
	utils.InitServer()
}

func main() {
	Initialize()
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Get("/pincode/:pincode?", func(c *fiber.Ctx) error {
		pincode := c.Params("pincode")
		fmt.Println("pincode", pincode)
		jsonData := utils.CheckPincode(pincode, c)
		return c.JSON(jsonData)
	})

	app.Listen(":3000")
}
