package main

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type FavoriteProduct struct {
	ProductID string `json:"product_id"`
}
type FavoriteProductBSON struct {
	ProductID primitive.ObjectID `json:"product_id" bson:"product_id"`
}

type AllFavorites struct {
	Favorites []FavoriteProductBSON `json:"favorites" bson:"favorites"`
}

type jwtClaims struct {
	UserID uuid.UUID `json:"user_id"`
	Email  string    `json:"email"`
	jwt.RegisteredClaims
}
