package main

import (
	"github.com/gin-gonic/gin"
)

func authRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		loggedIn := checkLoggedIn(c)
		if !loggedIn {
			c.JSON(401, gin.H{
				"message": "Authentication required. Please log in to access this resource.",
				"error":   "Authentication error",
			})
			c.Abort()
		}
		c.Next()
	}
}

func notAuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		loggedIn := checkLoggedIn(c)
		if loggedIn {
			c.JSON(401, gin.H{
				"message": "Unvalid request",
				"error":   "User has already logged in",
			})
			c.Abort()
		}
		c.Next()
	}
}

func checkLoggedIn(c *gin.Context) bool {
	_, err := getCookie(c)

	return err == nil
}
