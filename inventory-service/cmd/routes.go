package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func initRoutes(r *gin.Engine) {
	handlers := initHandlers()
	//cors
	r.Use(func(c *gin.Context) {
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
	})

	route := r.Group("/api/v1")

	route.GET("/:id", handlers.getStock)
	route.POST("/new", handlers.newProductStock)
	route.PATCH("/restock/:id", handlers.productReStock)
	route.PATCH("/cancel", handlers.cancelReservation)
	route.PATCH("/reserved/:id", handlers.updateStockViaReserved)
	route.PATCH("/sold/:id", handlers.updateStockViaSold)
}
