package main

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type UserBasicInfo struct {
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
type UserInfoDB struct {
	ID uuid.UUID `json:"user_id"`
	UserBasicInfo
	IsAdmin bool `json:"is_admin"`
}
type jwtClaims struct {
	UserID uuid.UUID `json:"user_id"`
	Email  string    `json:"email"`
	jwt.RegisteredClaims
}

type ChangePassword struct {
	CurrentPassword string `json:"current_password"`
	NewPassword     string `json:"new_password"`
}
