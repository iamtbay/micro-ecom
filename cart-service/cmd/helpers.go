package main

import (
	"errors"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

func isCookieValid(c *gin.Context) (uuid.UUID, error) {
	tokenString, err := c.Cookie("accessToken")
	if err != nil {
		return uuid.UUID{}, err
	}

	token, err := jwt.ParseWithClaims(tokenString, &jwtClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			return uuid.UUID{}, errors.New("invalid token, please login again")
		}
		return uuid.UUID{}, err
	}

	if claims, ok := token.Claims.(*jwtClaims); ok {
		return claims.UserID, nil
	} else {
		return uuid.UUID{}, err
	}

}
