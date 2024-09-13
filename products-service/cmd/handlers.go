package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProductsHandler struct{}

func initHandler() *ProductsHandler {
	return &ProductsHandler{}
}

var services = initServices()

// GET SINGLE PRODUCT
func (x *ProductsHandler) getSingleProduct(c *gin.Context) {
	id := c.Param("id")

	//get item
	productInfo, err := services.getSingleProduct(id)
	if err != nil {
		errJSON(http.StatusBadRequest, err, "Something went wrong", c)
	}

	//return
	c.JSON(200, gin.H{
		"message": "Succesfull",
		"data":    productInfo,
	})

}

// GET ALL PRODUCTS
func (x *ProductsHandler) getProducts(c *gin.Context) {
	page := c.Param("page")
	fmt.Println("page is ", page)

	products, err := services.getProducts()
	if err != nil {
		errJSON(http.StatusBadRequest, err, "something went wrong", c)
		return
	}

	c.JSON(200, gin.H{
		"message": "Succesfull",
		"data":    products,
	})

}

// ADD Product
func (x *ProductsHandler) addProduct(c *gin.Context) {
	var productInfo NewProduct

	//bind json
	err := c.BindJSON(&productInfo)
	if err != nil {
		errJSON(http.StatusInternalServerError, err, "Something went wrong while binding json", c)
		return
	}

	//service
	err = services.addProduct(&productInfo)
	if err != nil {
		errJSON(http.StatusInternalServerError, err, "Something went wrong while adding item", c)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Succesfully added",
		"data":    productInfo,
	})
}

// PATCH Product
func (x *ProductsHandler) editProduct(c *gin.Context) {
	var productInfo NewProduct

	//bind json
	err := c.BindJSON(&productInfo)
	if err != nil {
		errJSON(http.StatusInternalServerError, err, "Something went wrong while binding json", c)
		return
	}

	//get id
	id := c.Param("id")

	//service transaction
	newProductInfo, err := services.editProduct(id, &productInfo)
	if err != nil {
		errJSON(http.StatusInternalServerError, err, "Something went wrong while editing product", c)
		return
	}

	//return
	c.JSON(http.StatusOK, gin.H{
		"message": "succesfully edited",
		"data":    newProductInfo,
	})
}

// DELETE Product
func (x *ProductsHandler) deleteProduct(c *gin.Context) {
	id := c.Param("id")

	//delete product
	err := services.deleteProduct(id)
	if err != nil {
		errJSON(http.StatusInternalServerError, err, "Something went wrong while deleting product", c)
		return
	}
	//return
	c.JSON(http.StatusOK, gin.H{
		"message": "succesfully deleted",
		"data":    nil,
	})
}

// RETURN JSON
func errJSON(status int, err error, message string, c *gin.Context) {
	c.JSON(status, gin.H{
		"message": message,
		"error":   err.Error(),
	})
}
