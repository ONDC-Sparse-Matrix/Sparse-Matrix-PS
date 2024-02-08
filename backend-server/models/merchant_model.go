package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Merchant struct {
	Id       primitive.ObjectID `json:"id,omitempty"`
	Name     string             `json:"name,omnitempty" validate:"required"`
	Email    string             `json:"email,omnitempty" validate:"required"`
	Products string             `json:"products,omitempty" validate:"required"`
}
