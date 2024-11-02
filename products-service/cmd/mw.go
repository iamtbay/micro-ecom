package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func cookieCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString, err := c.Cookie("accessToken")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Authentication required. Please log in to access this resource.",
				"error":   "Authentication error",
			})
			c.Abort()
			return
		}
		//
		_, err = parseJWT(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthorized user",
				"message": err.Error(),
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
