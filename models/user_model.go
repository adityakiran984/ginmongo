package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	Id       primitive.ObjectID `json:"id"`
	Name     string             `json:"name,omitempty" validate:"required,min=5,max=100"`
	Location string             `json:"location,omitempty" validate:"required"`
	Title    string             `json:"title,omitempty" validate:"required, min=1, max=10"`
}
