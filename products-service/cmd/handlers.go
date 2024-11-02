package main

import (
	"fmt"
	"net/http"
	"strconv"

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
		errJSON(http.StatusBadRequest, err, "Something went wrong while getting product", c)
		return
	}

	//return
	c.JSON(200, gin.H{
		"message": "Product found succesfully",
		"data":    productInfo,
	})

}

// GET ALL PRODUCTS
func (x *ProductsHandler) getProducts(c *gin.Context) {
	//get page
	page := c.Query("page")
	pageInt, err := strconv.Atoi(page)
	if err != nil {
		errJSON(http.StatusBadRequest, err, "Something went wrong while page value converting ", c)
		return
	}

	//
	products, pageInfos, err := services.getProducts(int64(pageInt))
	if err != nil {

		errJSON(http.StatusBadRequest, err, "something went wrong while getting products", c)
		return
	}

	c.JSON(200, gin.H{
		"pages": gin.H{
			"page":                pageInfos.CurrentPage,
			"totalPage":           pageInfos.TotalPage,
			"total_product_count": pageInfos.TotalProductCount,
		},
		"message": fmt.Sprintf("%v products found", len(products)),
		"data":    products,
	})

}

// !
// ADD Product
func (x *ProductsHandler) addProduct(c *gin.Context) {
	var productInfo NewProduct
	//check cookie is valid?
	userID, _ := isCookieValid(c)

	//bind json
	err := c.BindJSON(&productInfo)
	productInfo.AddedBy = userID
	if err != nil {
		errJSON(http.StatusInternalServerError, err, "Something went wrong while reading product informations", c)
		return
	}

	//service
	err = services.addProduct(&productInfo)
	if err != nil {
		errJSON(http.StatusInternalServerError, err, "Something went wrong while adding product", c)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Product succesfully added",
		"data":    productInfo,
	})
}

// PATCH Product
func (x *ProductsHandler) editProduct(c *gin.Context) {
	var productInfo NewProduct
	//check cookie is valid?
	userID, _ := isCookieValid(c)
	
	//bind json
	err := c.BindJSON(&productInfo)
	if err != nil {
		errJSON(http.StatusInternalServerError, err, "Something went wrong while reading product informatins", c)
		return
	}

	//get id
	id := c.Param("id")

	productInfo.AddedBy = userID
	newProductInfo, err := services.editProduct(id, &productInfo)
	if err != nil {
		errJSON(http.StatusInternalServerError, err, "Something went wrong while editing product", c)
		return
	}

	//return
	c.JSON(http.StatusOK, gin.H{
		"message": "Product succesfully updated",
		"data":    newProductInfo,
	})
}

// !
// DELETE Product
func (x *ProductsHandler) deleteProduct(c *gin.Context) {
	id := c.Param("id")

	//check cookie is valid?
	userID, _ := isCookieValid(c)
	
	//delete product
	err := services.deleteProduct(id, userID)
	if err != nil {
		errJSON(http.StatusInternalServerError, err, "Something went wrong while deleting product", c)
		return
	}
	//return
	c.JSON(http.StatusOK, gin.H{
		"message": "Product succesfully deleted",
		"data":    nil,
	})
}

// RETURN ERROR JSON
func errJSON(status int, err error, message string, c *gin.Context) {
	c.JSON(status, gin.H{
		"message": message,
		"error":   err.Error(),
	})
}
