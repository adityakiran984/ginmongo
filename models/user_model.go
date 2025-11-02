package models

import "go.mongodb.org/mongo-driver/v2/bson"

type User struct {
	Id       bson.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name     string        `json:"name,omitempty" validate:"required,min=5,max=100" bson:"name"`
	Location string        `json:"location,omitempty" validate:"required" bson:"location"`
	Title    string        `json:"title,omitempty" validate:"required,min=1,max=10" bson:"title"`
}
