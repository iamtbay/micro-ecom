package main

import "github.com/gin-gonic/gin"

func initRoutes(r *gin.Engine) {
	route := r.Group("/api/v1")

	handler := initHandler()

	route.GET("/product/:id", handler.GetProductReviewsByProductID)
	route.GET("/:id", checkCookie(), handler.GetReviewByID)
	route.POST("/:id", checkCookie(), handler.NewReview)
	route.PATCH("/:id", checkCookie(), handler.EditReviewByReviewID)
	route.DELETE("/:id", checkCookie(), handler.DeleteReviewByReviewID)
}
