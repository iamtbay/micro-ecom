package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func initRoutes(r *gin.Engine) {
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

	handlers := initHandlers()
	route := r.Group("/api/v1")

	route.GET("/check", handlers.check)

	route.POST("/login", notAuthRequired(), handlers.login)
	route.POST("/signup", notAuthRequired(), handlers.signup)
	//
	route.POST("/logout", authRequired(), handlers.logout)
	route.PATCH("/edit", authRequired(), handlers.edit)
	route.PATCH("/changepassword", authRequired(), handlers.changePassword)
	route.DELETE("/delete", authRequired(), handlers.delete)

}
