package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handlers struct{}

func initHandlers() *Handlers {
	return &Handlers{}
}

var services = initServices()

//!
// NEW PRODUCT STOCK

func (x *Handlers) newProductStock(c *gin.Context) {
	var product Product

	err := c.BindJSON(&product)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = services.newProductStock(product)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusBadRequest, gin.H{"message": "Succesfully added!"})
}

// !
// GET PRODUCT STOCK BY ID
func (x *Handlers) getStock(c *gin.Context) {
	productID := c.Param("id")

	data, err := services.getStock(productID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"msg": "get stock", "data": data})
}

// !
// GET PRODUCT STOCK BY ID
func (x *Handlers) productReStock(c *gin.Context) {

	var product Product
	product.ProductID = c.Param("id")
	err := c.BindJSON(&product)
	if err != nil {
		c.JSON(200, gin.H{"error": err.Error()})
		return
	}
	err = services.productReStock(product)
	if err != nil {
		c.JSON(200, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"msg": "product re stock"})
}

// !
// UPDATE STOCK VIA PRODUCT RESERVED
func (x *Handlers) cancelReservation(c *gin.Context) {
	var product ProductData
	err := c.BindJSON(&product)
	if err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}

	err = services.cancelReservation(product)
	if err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}

	c.JSON(400, gin.H{"message": "succesfully cancelled"})
}

// !
// UPDATE STOCK VIA PRODUCT RESERVED
func (x *Handlers) updateStockViaReserved(c *gin.Context) {
	var product ProductData
	product.ProductID = c.Param("id")
	err := c.BindJSON(&product)
	if err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}
	err = services.updateStockViaReserved(product)
	if err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Succesfully reserved!"})

}

// !
// UPDATE STOCK VIA PRODUCT SOLD
func (x *Handlers) updateStockViaSold(c *gin.Context) {
	var product ProductData
	product.ProductID = c.Param("id")
	err := c.BindJSON(&product)
	if err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}
	err = services.updateStockViaSold(product)
	if err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Succesfully marked as sold!"})
}
