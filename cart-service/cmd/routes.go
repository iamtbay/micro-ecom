package main

import "github.com/gin-gonic/gin"

func initRoutes(r *gin.Engine) {
	handlers := initHandlers()
	route := r.Group("/api/v1")

	route.GET("/healthCheck", handlers.healthCheck)

	route.GET("/cart", handlers.getCart)
	route.POST("/cart/:id", handlers.addToCart)
	route.PATCH("/cart/:id", handlers.updateQuantityOfProduct)
	route.DELETE("/cart/:id", handlers.deleteProductOnCart)
	
	route.POST("/checkout", handlers.checkOut)
}
