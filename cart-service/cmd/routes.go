package main

import "github.com/gin-gonic/gin"

func initRoutes(r *gin.Engine) {
	handlers := initHandlers()
	route := r.Group("/api/v1")
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
			c.JSON(400, gin.H{"error": "Please login first!"})
			c.Abort()
			return
		}

		_, err = parseJWT(token)
		if err != nil {
			c.JSON(401, gin.H{"error": "Unauthorized user"})
		}
	}
}
