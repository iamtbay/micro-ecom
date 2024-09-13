package main

import "github.com/gin-gonic/gin"

func initRoutes(r *gin.Engine) {

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
