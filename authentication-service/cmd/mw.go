package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func authRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		loggedIn := checkLoggedIn(c)
		if !loggedIn {
			fmt.Println("unauthorized")
			c.JSON(401, gin.H{
				"message": "Unauthorized!",
			})
			c.Abort()
		}
		c.Next()
	}
}

//
func notAuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		loggedIn := checkLoggedIn(c)
		if loggedIn {
			fmt.Println("already logged in")
			c.JSON(401, gin.H{
				"message": "you can't use this!",
			})
			c.Abort()
		}
		c.Next()
	}
}

func checkLoggedIn(c *gin.Context) bool {
	_, err := getCookie(c)

	//parse jwt here
	return err == nil

}
