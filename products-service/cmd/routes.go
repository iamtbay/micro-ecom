package main

import "github.com/gin-gonic/gin"

func initRotues(r *gin.Engine) {
	route := r.Group("/api/v1/products")

	handlers := initHandler()

	route.GET("/s/:id", handlers.getSingleProduct)
	route.PATCH("/s/:id", handlers.editProduct)
	route.DELETE("/s/:id", handlers.deleteProduct)

	route.GET("/:page", handlers.getProducts)
	route.POST("/add", handlers.addProduct)
}
