package main

import (
	"fmt"
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
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": reviews})
}

// !
// GET PRODUCT REVIEWS BY PRODUCT ID
func (x *Handlers) GetReviewByID(c *gin.Context) {
	id := c.Param("id")
	reviews, err := services.getReviewByID(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": reviews})
}

// !
// GET PRODUCT REVIEWS BY PRODUCT ID
func (x *Handlers) NewReview(c *gin.Context) {
	productID := c.Param("id")
	userID, err := isCookieValid(c)
	fmt.Println(userID)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	var review NewReview
	err = c.BindJSON(&review)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	review.UserID = userID

	//servie
	err = services.newReview(review, productID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Review sent successfully",
	})
}

// !
// GET PRODUCT REVIEWS BY PRODUCT ID
func (x *Handlers) EditReviewByReviewID(c *gin.Context) {
	reviewID := c.Param("id")
	userID, err := isCookieValid(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	var review NewReview
	err = c.BindJSON(&review)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	review.UserID = userID
	err = services.editReviewByReviewID(reviewID, review)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Review updated successfully",
	})
}

// !
// GET PRODUCT REVIEWS BY PRODUCT ID
func (x *Handlers) DeleteReviewByReviewID(c *gin.Context) {
	reviewID := c.Param("id")
	userID, err := isCookieValid(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	err = services.deleteReviewByReviewID(reviewID, userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Review deleted succesfully",
	})

}
