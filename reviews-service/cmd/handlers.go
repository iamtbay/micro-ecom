package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handlers struct{}

func initHandler() *Handlers {
	return &Handlers{}
}

var services = initServices()

// !
// GET PRODUCT REVIEWS BY PRODUCT ID
func (x *Handlers) GetProductReviewsByProductID(c *gin.Context) {
	id := c.Param("id")

	reviews, err := services.getProductReviewsByProductID(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"message": "Error while getting product reviews",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":    reviews,
		"message": "Reviews got succesfully",
	})
}

// !
// GET PRODUCT REVIEWS BY PRODUCT ID
func (x *Handlers) GetReviewByID(c *gin.Context) {
	id := c.Param("id")
	reviews, err := services.getReviewByID(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"message": "Error while getting review",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":    reviews,
		"message": "Review got successfully",
	})
}

// !
// GET PRODUCT REVIEWS BY PRODUCT ID
func (x *Handlers) NewReview(c *gin.Context) {
	productID := c.Param("id")
	userID, _ := parseJWT(c)

	var review NewReview
	err := c.BindJSON(&review)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   err.Error(),
			"message": "Error while reading review informations",
		})
		return
	}
	review.UserID = userID

	//servie
	err = services.newReview(review, productID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"message": "Error while adding a review",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Review added successfully",
		"data":    nil,
	})
}

// !
// GET PRODUCT REVIEWS BY PRODUCT ID
func (x *Handlers) EditReviewByReviewID(c *gin.Context) {
	reviewID := c.Param("id")
	userID, _ := parseJWT(c)

	var review NewReview
	err := c.BindJSON(&review)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   err.Error(),
			"message": "Error while reading review informations",
		})
		return
	}
	review.UserID = userID
	err = services.editReviewByReviewID(reviewID, review)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"message": "Error while updating review",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Review updated successfully",
		"data":    nil,
	})
}

// !
// GET PRODUCT REVIEWS BY PRODUCT ID
func (x *Handlers) DeleteReviewByReviewID(c *gin.Context) {
	reviewID := c.Param("id")
	userID, _ := parseJWT(c)

	err := services.deleteReviewByReviewID(reviewID, userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"message": "Error while deleting review",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Review deleted succesfully",
		"data":    nil,
	})

}
