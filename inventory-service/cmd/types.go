package main

import (
	"encoding/json"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type Product struct {
	ID             uuid.UUID       `json:"id"`
	ProductID      string          `json:"product_id"`
	Properties     json.RawMessage `json:"properties"`
	AvailableStock int64           `json:"available_stock"`
	ReservedStock  int64           `json:"reserved_stock"`
}

type ProductData struct {
	ProductID string `json:"product_id"`
	Quantity  int    `json:"quantity"`
}

type jwtClaims struct {
	UserID uuid.UUID `json:"user_id"`
	Email  string    `json:"email"`
	jwt.RegisteredClaims
}
