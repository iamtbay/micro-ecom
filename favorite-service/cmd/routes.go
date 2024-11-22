package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func authMW() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, err := c.Cookie("accessToken")
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "You have to login!",
			})
			c.Abort()
		}
		c.Next()
	}
}

func initRoutes(r *gin.Engine) {

	route := r.Group("api/v1")
	handler := initHandlers()

	route.Use(authMW())
	route.GET("/favorites", handler.getAllFavorites)
	route.POST("/favorites/:id", handler.newFavorite)
	route.DELETE("/favorites/:id", handler.removeFavorite)

}
