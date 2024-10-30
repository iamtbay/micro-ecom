package main

import "github.com/gin-gonic/gin"

func initRoutes(r *gin.Engine) {
	handlers := initHandler()
	addressHandlers := initAdressesHandler()

	ordersRoute := r.Group("/api/v1")
	ordersRoute.Use(cookieRequired())
	ordersRoute.GET("/", handlers.healthCheck)
	ordersRoute.POST("/newOrder", handlers.newOrder)

	ordersRoute.GET("/:id", handlers.getSingleOrder)
	ordersRoute.GET("/user/:id", handlers.getOrdersByUserID)

	ordersRoute.DELETE("/:id", handlers.deleteOrder)

	//addresses routes
	addressRoute := r.Group("/api/v1/address")
	addressRoute.Use(cookieRequired())

	addressRoute.GET("", addressHandlers.GetAddresses)
	addressRoute.POST("", addressHandlers.AddNewAddress)

	addressRoute.GET("/:id", addressHandlers.GetSingleAddressByID)
	addressRoute.PATCH("/:id", addressHandlers.EditAddressByID)
	addressRoute.DELETE("/:id", addressHandlers.DeleteAddressByID)
}

func cookieRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := c.Cookie("accessToken")
		if err != nil {
			c.JSON(400, gin.H{"message": "Please login first!"})
			c.Abort()
			return
		}
		//
		_, err = parseJWT(token)
		if err != nil {
			c.JSON(401, gin.H{"message": "Unauthorized!"})
			c.Abort()
			return
		}

		c.Next()
	}
}
