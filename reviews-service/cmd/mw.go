package main

import (
	"errors"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

func corsMW() gin.HandlerFunc {
	return func(c *gin.Context) {
		origin := c.Request.Header.Get("Origin")
		if origin == "http://localhost:5173" || origin == "http://127.0.0.1:5173" || origin == "http://localhost:8080" {
			c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
		}

		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST,PATCH, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Authorization")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}
		c.Next()
	}
}

func checkCookie() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, err := c.Cookie("accessToken")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Authentication required. Please log in to access this resource.",
				"error":   "Authentication error",
			})
			c.Abort()
			return
		}
		_, err = parseJWT(c)
		if err != nil {
			c.JSON(401, gin.H{
				"error":   "Unauthorized user",
				"message": err.Error(),
			})
			c.Abort()
			return
		}
	}
}

func parseJWT(c *gin.Context) (uuid.UUID, error) {
	tokenString, _ := c.Cookie("accessToken")

	token, err := jwt.ParseWithClaims(tokenString, &jwtClaims{}, func(token *jwt.Token) (interface{}, error) {
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
		return uuid.UUID{}, errors.New("something went wrong while verifying user")
	}

}