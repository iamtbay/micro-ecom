package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type OrderHandler struct{}

func initOrderHandler() *OrderHandler {
	return &OrderHandler{}
}

// HEALTH CHECK
func (x *OrderHandler) healthCheck(c *gin.Context) {
	c.JSON(200, "health check")
}

// GET SINGLE ORDER
func (x *OrderHandler) getSingleOrder(c *gin.Context) {
	id := c.Param("id")
	fmt.Println("getSingleOrder", id)
	c.JSON(200, id)
}

// NEW ORDER
func (x *OrderHandler) newOrder(c *gin.Context) {
	fmt.Println("newOrder")
	c.JSON(200, "new order")
}

// EDIT ORDER
func (x *OrderHandler) editOrder(c *gin.Context) {
	id := c.Param("id")
	fmt.Println("editOrder", id)
	c.JSON(200, id)
}

// DELETE ORDER
func (x *OrderHandler) deleteOrder(c *gin.Context) {
	id := c.Param("id")
	fmt.Println("deleteOrder", id)
	c.JSON(200, id)
}
