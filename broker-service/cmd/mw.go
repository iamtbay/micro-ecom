package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func cookieRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		cookie := c.Request.Header.Get("cookie")
		if cookie == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "No access!"})
			c.Abort()
		}
		c.Next()
	}
}
