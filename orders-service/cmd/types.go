package main

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type jwtClaims struct {
	UserID uuid.UUID `json:"user_id"`
	Email  string    `json:"email"`
	jwt.RegisteredClaims
}

type Order struct {
	OrderID    uuid.UUID `json:"order_id,omitempty"`
	CustomerID uuid.UUID `json:"customer_id"`
	Products   []Product `json:"products,omitempty"`
	AddressID  uuid.UUID `json:"address_id"`
	TotalPrice float64   `json:"total_price"`
	OrderDate  time.Time `json:"order_date"`
	IsActive   bool      `json:"is_active"`
}

type Product struct {
	ID       string  `json:"_id"`
	Quantity int     `json:"quantity"`
	Price    float64 `json:"price"`
}

type ShippingAdress struct {
	StreetAdress1 string `json:"street_adress_1"`
	StreetAdress2 string `json:"street_adress_2"`
	City          string `json:"city"`
	State         string `json:"state"`
	PostalCode    string `json:"postal_code"`
	Country       string `json:"country"`
}

type NewAddress struct {
	UserID      uuid.UUID `json:"user_id"`
	AddressName string    `json:"address_name"`
	Street      string    `json:"street"`
	City        string    `json:"city"`
	State       string    `json:"state"`
	PostalCode  string    `json:"postal_code"`
	Country     string    `json:"country"`
}

type GetAddresses struct {
	ID          uuid.UUID `json:"id"`
	UserID      uuid.UUID `json:"user_id"`
	AddressName string    `json:"address_name"`
	Street      string    `json:"street"`
	City        string    `json:"city"`
	State       string    `json:"state"`
	PostalCode  string    `json:"postal_code"`
	Country     string    `json:"country"`
	IsDeleted   bool      `json:"-"`
}

type MessageType struct {
	Message string `json:"message"`
}
