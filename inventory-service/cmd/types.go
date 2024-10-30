package main

import (
	"encoding/json"

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
