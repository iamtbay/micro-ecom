package main

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type NewProduct struct {
	Name    string `json:"name" bson:"name"`
	Brand   string `json:"brand" bson:"brand"`
	Content string `json:"content" bson:"content"`
	AddedBy string `json:"added_by" bson:"added_by"`
}

type GetProduct struct {
	ID      primitive.ObjectID `json:"_id,omitempty" bson:"_id"`
	Name    string             `json:"name" bson:"name"`
	Brand   string             `json:"brand" bson:"brand"`
	Content string             `json:"content" bson:"content"`
	AddedBy string             `json:"added_by" bson:"added_by"`
}
