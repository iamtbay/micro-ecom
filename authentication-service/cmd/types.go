package main

import "github.com/golang-jwt/jwt/v5"

type UserBasicInfo struct {
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
type UserInfoDB struct {
	ID int32 `json:"user_id"` //change to google id
	UserBasicInfo
}
type jwtClaims struct {
	UserID int32    `json:"user_id"`
	Email  string `json:"email"`
	jwt.RegisteredClaims
}

type NewPassword struct {
	Password string `json:"new_password"`
}