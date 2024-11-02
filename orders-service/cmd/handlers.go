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
			"message": "Error while getting orders",
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
			"error":   err.Error(),
			"message": "Error while getting order",
		})
		return
	}

	//return response
	c.JSON(200, gin.H{
		"message": "Order found",
		"data":    data,
	})
}

// !
// NEW ORDER
func (x *Handler) newOrder(c *gin.Context) {
	var order Order
	err := c.BindJSON(&order)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"message": "Error while reading order informations",
		})
		return
	}

	//check user
	userID, _ := isCookieValid(c)

	order.CustomerID = userID
	//services
	data, err := services.newOrder(order)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"message": "Error while creating a new order",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Ordered successfully",
		"data":    data,
	})
}

// !
// DELETE ORDER
func (x *Handler) deleteOrder(c *gin.Context) {
	orderID := c.Param("id")
	userID, _ := isCookieValid(c)

	err := services.deleteOrder(orderID, userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"message": "Error while deleting order",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Order succesfully deleted",
	})
}
