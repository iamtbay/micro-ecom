package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func initRoutes(r *gin.Engine) {
	handlers := initHandlers()
	route := r.Group("/api/v1")
	route.Use(corsMW())
	route.Use(cookieRequired())

	route.GET("/healthCheck", handlers.healthCheck)

	route.GET("/cart", handlers.getCart)
	route.POST("/cart/:id", handlers.addToCart)
	route.PATCH("/cart/:id", handlers.updateQuantityOfProduct)
	route.DELETE("/cart/:id", handlers.deleteProductOnCart)

	route.POST("/checkout", handlers.checkOut)
}

func cookieRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := c.Cookie("accessToken")
		if err != nil {
			c.JSON(400, gin.H{
				"message": "Authentication required. Please log in to access this resource.",
				"error":   "Authentication error",
			})
			c.Abort()
			return
		}

		_, err = parseJWT(token)
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
