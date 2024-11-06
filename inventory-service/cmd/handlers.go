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
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"message": "Error while reading product informations",
		})
		return
	}

	err = services.newProductStock(product)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"message": "Error while adding stock",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Succesfully added!", "data": nil})
}

// !
// GET PRODUCT STOCK BY ID
func (x *Handlers) getProductStock(c *gin.Context) {
	productID := c.Param("id")

	data, err := services.getStock(productID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"message": "Error while getting stock",
		})
		return
	}

	c.JSON(200, gin.H{"message": "Stock got succesfully", "data": data})
}

// !
// GET PRODUCT STOCK BY ID
func (x *Handlers) restockProduct(c *gin.Context) {

	var product Product
	product.ProductID = c.Param("id")
	err := c.BindJSON(&product)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"message": "Error while reading product informations",
		})
		return
	}
	err = services.productReStock(product)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"message": "Error while adding product stock",
		})
		return
	}

	c.JSON(200, gin.H{"message": "Stock is updated"})
}

// !
// UPDATE STOCK VIA PRODUCT RESERVED
func (x *Handlers) cancelStockReservation(c *gin.Context) {
	var product ProductData
	err := c.BindJSON(&product)
	product.ProductID = c.Param("id")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"message": "Error while reading product informations",
		})
		return
	}

	err = services.cancelReservation(product)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"message": "Error while cancelling a reservation",
		})
		return

	}

	c.JSON(200, gin.H{"message": "Product successfully cancelled from reservation queue"})
}

// !
// UPDATE STOCK VIA PRODUCT RESERVED
func (x *Handlers) confirmStockReservation(c *gin.Context) {
	var product ProductData
	product.ProductID = c.Param("id")
	err := c.BindJSON(&product)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"message": "Error while reading product informations",
		})
		return

	}
	err = services.updateStockViaReserved(product)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"message": "Error while a product registering on reservation queue",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Succesfully reserved!"})
}

// !
// UPDATE STOCK VIA PRODUCT SOLD
func (x *Handlers) updateStockAfterSale(c *gin.Context) {
	var product ProductData
	product.ProductID = c.Param("id")
	err := c.BindJSON(&product)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"message": "Error while reading product informations",
		})
		return
	}
	err = services.updateStockViaSold(product)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"message": "Error while product updating as sold",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Succesfully marked as sold!"})
}
