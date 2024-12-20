package handlersPackage

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

type AuthChangePassword struct {
	CurrentPassword string `json:"current_password"`
	NewPassword     string `json:"new_password"`
}

type AuthRequest struct {
	Email       string `json:"email"`
	Name        string `json:"name"`
	Surname     string `json:"surname"`
	Password    string `json:"password,omitempty"`
	NewPassword string `json:"new_password,omitempty"`
}

type ProductRequest struct {
}

type ProductData struct {
	Name     string    `json:"name"`
	Brand    string    `json:"brand"`
	Content  string    `json:"content"`
	Price    float64   `json:"price"`
	Stock    int       `json:"stock"`
	Images   []string  `json:"images"`
	Added_By uuid.UUID `json:"added_by"`
}

type CartRequest struct {
	Name     string  `json:"name"`
	Quantity int     `json:"quantity"`
	Price    float64 `json:"price"`
}

type OrderRequest struct {
	OrderID    uuid.UUID      `json:"order_id,omitempty"`
	CustomerID uuid.UUID      `json:"customer_id"`
	Products   []OrderProduct `json:"products,omitempty"`
	AddressID  uuid.UUID      `json:"address_id"`
	TotalPrice float64        `json:"total_price"`
	OrderDate  time.Time      `json:"order_date"`
	IsActive   bool           `json:"is_active"`
}

type OrderProduct struct {
	ID       string  `json:"_id"`
	Quantity int     `json:"quantity"`
	Price    float64 `json:"price"`
}

type CartQuantityRequest struct {
	SetExact bool `json:"set_exact"`
}
type CartCheckOutType struct {
	AddressID uuid.UUID `json:"address_id"`
}

type NewAddress struct {
	AddressName string `json:"address_name"`
	Street      string `json:"street"`
	City        string `json:"city"`
	State       string `json:"state"`
	PostalCode  string `json:"postal_code"`
	Country     string `json:"country"`
}

type NewReview struct {
	Name    string `json:"name" bson:"name"`
	Surname string `json:"surname" bson:"surname"`
	Point   int64  `json:"point" bson:"point"`
	Comment string `json:"comment" bson:"comment"`
}

type InventoryProduct struct {
	ID             uuid.UUID       `json:"id"`
	ProductID      string          `json:"product_id"`
	Properties     json.RawMessage `json:"properties"`
	AvailableStock int64           `json:"available_stock"`
	ReservedStock  int64           `json:"reserved_stock"`
}

type InventoryProductSale struct {
	ProductID string `json:"product_id"`
	Quantity  int    `json:"quantity"`
}
