package controllers

import (
	"context"
	"backend-server/configs"
	"backend-server/models"
	"backend-server/responses"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var mapCollection *mongo.Collection = configs.GetCollection(configs.DB, "map")
var merchantsCollection *mongo.Collection = configs.GetCollection(configs.DB, "merchants")

func GetMerchants(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	pinCode := c.Params("pinCode")
	var merchants models.Map
	defer cancel()

	err := mapCollection.FindOne(ctx, bson.M{"pinCode": pinCode}).Decode(&merchants)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.Response{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	arr := merchants.Array
	length := len(arr)
	response := make([]models.Merchant, 0)
	for i := 0; i < length; i++ {
		var merchant models.Merchant
		objId, _ := primitive.ObjectIDFromHex(arr[i])
		err := merchantsCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&merchant)
		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(responses.Response{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
		} else {
			response = append(response, merchant)
		}
	}
	return c.Status(http.StatusOK).JSON(responses.Response{Status: http.StatusOK, Message: "success", Data: &fiber.Map{"data": response}})
}


func AddMerchants(c *fiber.Ctx) error {
	//! handle the stream and put data in database
	return nil
}