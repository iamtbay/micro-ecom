package main

import "github.com/gin-gonic/gin"

func initRoutes(r *gin.Engine) {
	route := r.Group("/api/v1")

	handler := initHandler()

	route.GET("/product/:id", handler.GetProductReviewsByProductID)
	route.GET("/:id", handler.GetReviewByID)
	route.POST("/:id", handler.NewReview)
	route.PATCH("/:id", handler.EditReviewByReviewID)
	route.DELETE("/:id", handler.DeleteReviewByReviewID)
}
