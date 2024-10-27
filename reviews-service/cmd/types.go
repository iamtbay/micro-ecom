package main

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	uuid "github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type NewReview struct {
	ProductID primitive.ObjectID `json:"product_id" bson:"product_id"`
	UserID    uuid.UUID          `json:"user_id" bson:"user_id"`
	Name      string             `json:"name" bson:"name"`
	Surname   string             `json:"surname" bson:"surname"`
	Point     int64              `json:"point" bson:"point"`
	Comment   string             `json:"comment" bson:"comment"`
	Date      time.Time          `json:"date" bson:"date"`
	IsDeleted bool               `json:"is_deleted" bson:"is_deleted"`
}
type NewReviewBSON struct {
	ProductID primitive.ObjectID `json:"product_id" bson:"product_id"`
	UserID    primitive.Binary   `json:"user_id" bson:"user_id"`
	Name      string             `json:"name" bson:"name"`
	Surname   string             `json:"surname" bson:"surname"`
	Point     int64              `json:"point" bson:"point"`
	Comment   string             `json:"comment" bson:"comment"`
	Date      time.Time          `json:"date" bson:"date"`
	IsDeleted bool               `json:"is_deleted" bson:"is_deleted"`
}
type GetReview struct {
	ReviewID  primitive.ObjectID `json:"_id" bson:"_id"`
	ProductID primitive.ObjectID `json:"product_id" bson:"product_id"`
	UserID    uuid.UUID          `json:"user_id" bson:"user_id"`
	Name      string             `json:"name" bson:"name"`
	Surname   string             `json:"surname" bson:"surname"`
	Point     int64              `json:"point" bson:"point"`
	Comment   string             `json:"comment" bson:"comment"`
	Date      time.Time          `json:"date" bson:"date"`
	IsDeleted bool               `json:"is_deleted" bson:"is_deleted"`
}
type GetReviewBSON struct {
	ReviewID  primitive.ObjectID `json:"_id" bson:"_id"`
	ProductID primitive.ObjectID `json:"product_id" bson:"product_id"`
	UserID    primitive.Binary   `json:"user_id" bson:"user_id"`
	Name      string             `json:"name" bson:"name"`
	Surname   string             `json:"surname" bson:"surname"`
	Point     int64              `json:"point" bson:"point"`
	Comment   string             `json:"comment" bson:"comment"`
	Date      time.Time          `json:"date" bson:"date"`
	IsDeleted bool               `json:"is_deleted" bson:"is_deleted"`
}

type jwtClaims struct {
	UserID uuid.UUID `json:"user_id"`
	Email  string    `json:"email"`
	jwt.RegisteredClaims
}
