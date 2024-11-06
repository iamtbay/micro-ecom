package main

import "github.com/gin-gonic/gin"

var handlers = initHandlers()

//
func initRoutes(r *gin.Engine) {
	route := r.Group("/api/v1")
	route.GET("/search", handlers.SearchProduct)
}
