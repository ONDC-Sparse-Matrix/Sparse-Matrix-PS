package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Map struct {
	Id       primitive.ObjectID   `json:"id,omitempty"`
	PIN_CODE string               `json:"pinCode,omitempty" validate:"required"`
	Array    []string `json:"merchants,omitempty" bson:"merchants,omitempty"`
}
