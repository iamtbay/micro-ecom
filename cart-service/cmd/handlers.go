package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Handlers struct{}

func initHandlers() *Handlers {
	return &Handlers{}
}

var services = initServices()

// !
// CHECK OUT
func (x *Handlers) checkOut(c *gin.Context) {
	userID, err := isCookieValid(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	var addressID CheckOutType
	//user has to select valid address
	err = c.BindJSON(&addressID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = services.checkOut(userID, addressID.AddressID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "ordered succesfully",
	})
}

// !
// HEALTH CHECK
func (x *Handlers) healthCheck(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "health check",
	})
}

// !
// GET CART
func (x *Handlers) getCart(c *gin.Context) {
	//check cookie
	userID, err := isCookieValid(c)
	if err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{
			"error": err.Error(),
		})
		return
	}

	//get cart
	data, err := services.getCart(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	//
	if len(data.Products) < 1 {
		c.JSON(http.StatusOK, gin.H{
			"message": "User cart is empty",
		})
		return
	}

	//return response
	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("%v items on cart", len(data.Products)),
		"data":    data,
	})

}

// !
// ADD TO CART
func (x *Handlers) addToCart(c *gin.Context) {
	productID := c.Param("id")
	//check cookie
	userID, err := isCookieValid(c)
	if err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{
			"error": err.Error(),
		})
		return
	}

	//get product as a json?
	var product CartItem
	err = c.BindJSON(&product)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	product.ProductID, err = primitive.ObjectIDFromHex(productID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	//service
	err = services.addToCart(userID, product)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return

	}

	//return
	c.JSON(http.StatusOK, gin.H{
		"message": "product added your cart.",
	})
}

// !
// UPDATE QUANTITY
func (x *Handlers) updateQuantityOfProduct(c *gin.Context) {
	quantity := c.Query("quantity")

	productID := c.Param("id")

	userID, err := isCookieValid(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		return
	}

	var isExact SetExact
	err = c.BindJSON(&isExact)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		return
	}
	msg, err := services.updateQuantityOfProduct(userID, productID, quantity, isExact.SetExact)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		return
	}

	//return response
	c.JSON(http.StatusOK, gin.H{
		"message": msg,
	})

}

// !
// DELETE CART
func (x *Handlers) deleteProductOnCart(c *gin.Context) {
	// check cookie
	userID, err := isCookieValid(c)
	if err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{
			"error": err.Error(),
		})
		return
	}
	// check item id
	productID := c.Param("id")

	// what changed?
	err = services.deleteProductOnCart(userID, productID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Product succesfully removed on your cart.",
	})
}
