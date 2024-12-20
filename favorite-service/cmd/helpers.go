package main

import (
	"errors"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

func isCookieValid(c *gin.Context) (uuid.UUID, error) {
	tokenString, _ := c.Cookie("accessToken")
	userID, err := parseJWT(tokenString)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		c.Abort()
	}
	return userID, nil

}

func parseJWT(tokenString string) (uuid.UUID, error) {
	token, err := jwt.ParseWithClaims(tokenString, &jwtClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			return uuid.UUID{}, errors.New("Token expired, please login again!")
		}
		return uuid.UUID{}, err
	}

	if claims, ok := token.Claims.(*jwtClaims); ok {
		return claims.UserID, nil
	} else {
		return uuid.UUID{}, err
	}
}
