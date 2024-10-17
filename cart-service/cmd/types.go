package main

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type jwtClaims struct {
	UserID uuid.UUID `json:"user_id"`
	Email  string    `json:"email"`
	jwt.RegisteredClaims
}

type CartItem struct {
	ProductID primitive.ObjectID `json:"_id"`
	Name      string             `json:"name"`
	Quantity  int                `json:"quantity"`
	Price     float64            `json:"price"`
}

type SetExact struct {
	SetExact bool `json:"set_exact"`
}

type CartOrder struct {
	UserID     uuid.UUID  `json:"customer_id"`
	AddressID  uuid.UUID  `json:"address_id"`
	Products   []CartItem `json:"products"`
	TotalPrice float64    `json:"total_price"`
}

type CheckOutType struct {
	AddressID uuid.UUID `json:"address_id"`
}

type UpdateProductType struct {
	ID    primitive.ObjectID `json:"_id"`
	Name  string             `json:"name"`
	Price float64            `json:"price"`
}
