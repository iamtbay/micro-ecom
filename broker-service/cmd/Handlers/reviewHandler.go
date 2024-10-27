package handlersPackage

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type ReviewHandler struct {
}

func InitReviewHandlers() *ReviewHandler {
	return &ReviewHandler{}
}

// !
// GET PRODUCT REVIEWS BY PROD ID
func (x *ReviewHandler) GetProductReviewsByProductID(c *gin.Context) {
	prodID := c.Param("id")
	req, _ := http.NewRequest("GET", "", nil)

	serviceResp, err := forwardRequest(fmt.Sprintf("%v/product/%v", os.Getenv("REVIEWS_SERVICE_URL"), prodID), req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	resp, err := io.ReadAll(serviceResp.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	defer serviceResp.Body.Close()
	c.Data(serviceResp.StatusCode, serviceResp.Header.Get("Content-Type"), resp)
}

// !
// GET REVIEW BY ID
func (x *ReviewHandler) GetReviewByID(c *gin.Context) {
	reviewID := c.Param("id")
	req, _ := http.NewRequest("GET", "", nil)

	serviceResp, err := forwardRequest(fmt.Sprintf("%v/%v", os.Getenv("REVIEWS_SERVICE_URL"), reviewID), req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	resp, err := io.ReadAll(serviceResp.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	defer serviceResp.Body.Close()
	c.Data(serviceResp.StatusCode, serviceResp.Header.Get("Content-Type"), resp)
}

// !
// NEW REVIEW
func (x *ReviewHandler) NewReview(c *gin.Context) {
	cookie := c.Request.Header.Get("cookie")
	prodID := c.Param("id")

	var review NewReview
	err := c.BindJSON(&review)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	jsonData, err := json.Marshal(review)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	req, _ := http.NewRequest("POST", "", bytes.NewBuffer(jsonData))
	req.Header.Add("cookie", cookie)
	serviceResp, err := forwardRequest(fmt.Sprintf("%v/%v", os.Getenv("REVIEWS_SERVICE_URL"), prodID), req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	defer serviceResp.Body.Close()

	resp, err := io.ReadAll(serviceResp.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.Data(serviceResp.StatusCode, serviceResp.Header.Get("Content-Type"), resp)
}

// !
// EDIT REVIEW
func (x *ReviewHandler) EditReviewByReviewID(c *gin.Context) {
	cookie := c.Request.Header.Get("cookie")
	reviewID := c.Param("id")

	var review NewReview
	err := c.BindJSON(&review)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	jsonData, err := json.Marshal(review)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	req, _ := http.NewRequest("PATCH", "", bytes.NewBuffer(jsonData))
	req.Header.Add("cookie", cookie)
	serviceResp, err := forwardRequest(fmt.Sprintf("%v/%v", os.Getenv("REVIEWS_SERVICE_URL"), reviewID), req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	defer serviceResp.Body.Close()

	resp, err := io.ReadAll(serviceResp.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.Data(serviceResp.StatusCode, serviceResp.Header.Get("Content-Type"), resp)
}

// !
// DELETE REVIEW
func (x *ReviewHandler) DeleteReviewByReviewID(c *gin.Context) {
	cookie := c.Request.Header.Get("cookie")
	reviewID := c.Param("id")

	req, _ := http.NewRequest("DELETE", "", nil)
	req.Header.Add("cookie", cookie)
	serviceResp, err := forwardRequest(fmt.Sprintf("%v/%v", os.Getenv("REVIEWS_SERVICE_URL"), reviewID), req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	defer serviceResp.Body.Close()

	resp, err := io.ReadAll(serviceResp.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.Data(serviceResp.StatusCode, serviceResp.Header.Get("Content-Type"), resp)
}
