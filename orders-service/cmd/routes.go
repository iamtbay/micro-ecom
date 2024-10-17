package main

import "github.com/gin-gonic/gin"

func initRoutes(r *gin.Engine) {
	handlers := initHandler()
	addressHandlers := initAdressesHandler()

	ordersRoute := r.Group("/api/v1")

	ordersRoute.GET("/", handlers.healthCheck)
	ordersRoute.POST("/newOrder", handlers.newOrder)

	ordersRoute.GET("/:id", handlers.getSingleOrder)
	ordersRoute.GET("/user/:id", handlers.getOrdersByUserID)

	ordersRoute.DELETE("/:id", handlers.deleteOrder)

	//addresses routes
	addressRoute := r.Group("/api/v1/address")

	addressRoute.GET("", addressHandlers.GetAddresses)
	addressRoute.POST("", addressHandlers.AddNewAddress)

	addressRoute.GET("/:id", addressHandlers.GetSingleAddressByID)
	addressRoute.PATCH("/:id", addressHandlers.EditAddressByID)
	addressRoute.DELETE("/:id", addressHandlers.DeleteAddressByID)
}
