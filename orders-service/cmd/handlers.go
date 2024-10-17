package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct{}

func initHandler() *Handler {
	return &Handler{}
}

var services = initServices()

// HEALTH CHECK
func (x *Handler) healthCheck(c *gin.Context) {
	c.JSON(200, "health check")
}

// !
// GET ORDERS BY USER ID
func (x *Handler) getOrdersByUserID(c *gin.Context) {
	userID := c.Param("id")

	//
	orders, err := services.getAllOrdersByUserID(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	//return response
	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("%v orders founded", len(orders)),
		"data":    orders,
	})
}

// !
// GET SINGLE ORDER
func (x *Handler) getSingleOrder(c *gin.Context) {
	id := c.Param("id")
	//
	data, err := services.getSingleOrder(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	//return response
	c.JSON(200, data)
}

// !
// NEW ORDER
func (x *Handler) newOrder(c *gin.Context) {
	var order Order
	err := c.BindJSON(&order)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	//check user
	userID, err := isCookieValid(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	order.CustomerID = userID
	//services
	data, err := services.newOrder(order)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "succesfully ordered",
		"data":    data,
	})
}

// !
// DELETE ORDER
func (x *Handler) deleteOrder(c *gin.Context) {
	orderID := c.Param("id")
	userID, err := isCookieValid(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		return
	}
	err = services.deleteOrder(orderID, userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Succesfully deleted",
	})
}
