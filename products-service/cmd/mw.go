package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func cookieCheck() gin.HandlerFunc {
	return func(c *gin.Context) {

		_, err := c.Cookie("accessToken")
		if err != nil {
			fmt.Println(err,"no cookie")
			c.Abort()
		}
		c.Next()
	}
}
