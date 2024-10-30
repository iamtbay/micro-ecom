package main

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type NewProduct struct {
	Name    string    `json:"name" bson:"name"`
	Brand   string    `json:"brand" bson:"brand"`
	Content string    `json:"content" bson:"content"`
	Price   float64   `json:"price" bson:"price"`
	Stock   int       `json:"stock"`
	AddedBy uuid.UUID `json:"added_by,omitempty" bson:"added_by"`
}
type NewProductBSON struct {
	Name    string           `json:"name" bson:"name"`
	Brand   string           `json:"brand" bson:"brand"`
	Content string           `json:"content" bson:"content"`
	Price   float64          `json:"price" bson:"price"`
	AddedBy primitive.Binary `json:"added_by,omitempty" bson:"added_by"`
}

type ProductInventoryType struct {
	ProductID string `json:"product_id"`
	Quantity  int    `json:"quantity"`
}

type GetProduct struct {
	ID      primitive.ObjectID `json:"_id,omitempty" bson:"_id"`
	Name    string             `json:"name" bson:"name"`
	Brand   string             `json:"brand" bson:"brand"`
	Content string             `json:"content" bson:"content"`
	Price   float64            `json:"price" bson:"price"`
	AddedBy uuid.UUID          `json:"added_by" bson:"added_by"`
}
type GetProductBSON struct {
	ID      primitive.ObjectID `json:"_id,omitempty" bson:"_id"`
	Name    string             `json:"name" bson:"name"`
	Brand   string             `json:"brand" bson:"brand"`
	Content string             `json:"content" bson:"content"`
	Price   float64            `json:"price" bson:"price"`
	AddedBy primitive.Binary   `json:"added_by" bson:"added_by"`
}

type EditProduct struct {
	Name    string  `json:"name" bson:"name"`
	Brand   string  `json:"brand" bson:"brand"`
	Content string  `json:"content" bson:"content"`
	Price   float64 `json:"price" bson:"price"`
}

type PageInfo struct {
	CurrentPage       int `json:"current_page"`
	TotalPage         int `json:"total_page"`
	TotalProductCount int `json:"total_product_count"`
}

type jwtClaims struct {
	UserID uuid.UUID `json:"user_id"`
	Email  string    `json:"email"`
	jwt.RegisteredClaims
}
