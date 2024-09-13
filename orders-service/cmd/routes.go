package main

import "github.com/gin-gonic/gin"

func initRoutes(r *gin.Engine) {
	handlers := initOrderHandler()

	ordersRoute := r.Group("/api/v1/orders")

	ordersRoute.GET("/", handlers.healthCheck)
	ordersRoute.POST("/newOrder", handlers.newOrder)

	ordersRoute.GET("/:id", handlers.getSingleOrder)
	ordersRoute.PATCH("/:id", handlers.editOrder)
	ordersRoute.DELETE("/:id", handlers.deleteOrder)
}
