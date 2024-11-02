package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func initRoutes(r *gin.Engine) {

	//cors
	r.Use(corsMW())

	route := r.Group("/api/v1")

	handlers := initHandler()

	route.GET("/s/:id", handlers.getSingleProduct)
	route.PATCH("/s/:id", cookieCheck(), handlers.editProduct)
	route.DELETE("/s/:id", cookieCheck(), handlers.deleteProduct)

	route.GET("", handlers.getProducts)
	route.POST("/add", cookieCheck(), handlers.addProduct)
}

func corsMW() gin.HandlerFunc {
	return func(c *gin.Context) {
		origin := c.Request.Header.Get("Origin")
		if origin == "http://localhost:5173" || origin == "http://127.0.0.1:5173" {
			c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
		}

		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST,PATCH, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Authorization")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			// Preflight OPTIONS talepleri için yanıt döndür
			c.AbortWithStatus(http.StatusNoContent)
			return
		}
		c.Next()
	}
}
