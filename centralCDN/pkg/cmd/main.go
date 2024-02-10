package main

import (
	"centralCDN/pkg/utils"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func Initialise() {
	utils.InitCacheServerList()
	utils.InitPincode()
	utils.InitServerRangeList()
}

func main() {
	app := fiber.New()
	Initialise()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Get("/pincode/:pincode?", func(c *fiber.Ctx) error {

		pincode := c.Params("pincode")

		fmt.Println("Pincode", pincode)
		body := utils.FetchMerchantData(pincode)
		// pincode := c.Params("pincode")
		// jsonData := []byte(`[
		// 	{"pincode" : "22114097", "merchantList" : ["merchant1", "merchant2", "merchant3"]},
		// 	{"pincode" : "22114098", "merchantList" : ["merchant1", "merchant2", "merchant3"]},
		// 	{"pincode" : "22114099", "merchantList" : ["merchant1", "merchant2", "merchant3"]},
		// 	{"pincode" : "22114096", "merchantList" : ["merchant1", "merchant2", "merchant3"]},
		// 	{"pincode" : "22114095", "merchantList" : ["merchant1", "merchant2", "merchant3"]},
		// 	{"pincode" : "22114094", "merchantList" : ["merchant1", "merchant2", "merchant3"]},
		// 	{"pincode" : "22114093", "merchantList" : ["merchant1", "merchant2", "merchant3"]}
		// ]`)

		// go utils.FetchMerchantData(pincode)
		return c.SendString(body)
	})

	app.Listen(":3001")
}

// /merchant/pincode/:pincode
