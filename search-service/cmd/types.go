package main

import (
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Product struct {
	ID      primitive.ObjectID `json:"_id,omitempty" bson:"_id"`
	Name    string             `json:"name" bson:"name"`
	Brand   string             `json:"brand" bson:"brand"`
	Content string             `json:"content" bson:"content"`
	Price   float64            `json:"price" bson:"price"`
	AddedBy uuid.UUID          `json:"added_by" bson:"added_by"`
}
type ProductIndex struct {
	Name    string    `json:"name" bson:"name"`
	Brand   string    `json:"brand" bson:"brand"`
	Content string    `json:"content" bson:"content"`
	Price   float64   `json:"price" bson:"price"`
	AddedBy uuid.UUID `json:"added_by" bson:"added_by"`
}
