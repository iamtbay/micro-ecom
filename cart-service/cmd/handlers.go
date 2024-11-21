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
	userID, _ := isCookieValid(c)

	var addressID CheckOutType
	//user has to select valid address
	err := c.BindJSON(&addressID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"message": "Error while reading user informations",
		})
		return
	}

	err = services.checkOut(userID, addressID.AddressID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"message": "Error on checkout",
		})
		return
	}

	c.JSON(http.StatusNoContent, gin.H{
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
	userID, _ := isCookieValid(c)

	//get cart
	data, err := services.getCart(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"message": "Error while getting users' cart",
		})
		return
	}

	//
	if len(data.Products) < 1 {
		c.JSON(http.StatusOK, gin.H{
			"message": "User cart is empty",
			"data":    nil,
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
	userID, _ := isCookieValid(c)

	//get product as a json?
	var product CartItem
	err := c.BindJSON(&product)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"message": "Error while reading user informations",
		})
		return
	}
	product.ProductID, err = primitive.ObjectIDFromHex(productID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"message": "Error on product process",
		})
		return
	}

	//service
	err = services.addToCart(userID, product)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"message": "Error while product adding to cart",
		})
		return

	}

	//return
	c.JSON(http.StatusNoContent, gin.H{
		"message": "product added your cart.",
	})
}

// !
// UPDATE QUANTITY
func (x *Handlers) updateQuantityOfProduct(c *gin.Context) {
	quantity := c.Query("quantity")

	productID := c.Param("id")

	userID, _ := isCookieValid(c)

	var isExact SetExact
	err := c.BindJSON(&isExact)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"message": "Error while reading user informations",
		})
		return
	}
	msg, err := services.updateQuantityOfProduct(userID, productID, quantity, isExact.SetExact)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
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
	userID, _ := isCookieValid(c)

	// check item id
	productID := c.Param("id")

	// what changed?
	err := services.deleteProductOnCart(userID, productID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"message": "Error while deleting product from cart",
		})
		return
	}

	c.JSON(http.StatusNoContent, gin.H{
		"message": "Product succesfully removed on your cart.",
	})
}
