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

type CartHandler struct{}

func InitCartHandlers() *CartHandler {
	return &CartHandler{}
}

// !
// GET CART
func (x *CartHandler) GetCart(c *gin.Context) {
	cookie := c.Request.Header.Get("cookie")

	//get jsons and marshal it
	req, err := http.NewRequest("GET", "", nil)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	req.Header.Add("cookie", cookie)

	serviceResp, err := forwardRequest(fmt.Sprintf("%v/cart", os.Getenv("CART_SERVICE_URL")), req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	defer serviceResp.Body.Close()

	respBody, err := io.ReadAll(serviceResp.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "message": "failed to read response from service"})
		return
	}
	c.Data(serviceResp.StatusCode, serviceResp.Header.Get("Content-Type"), respBody)
}

// !
// ADD TO CART
func (x *CartHandler) AddToCart(c *gin.Context) {
	productID := c.Param("id")
	//cookie
	cookie := c.Request.Header.Get("cookie")

	//get product data to json
	var product CartRequest
	err := c.BindJSON(&product)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	//marshal
	jsonData, err := json.Marshal(product)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	//create req
	req, err := http.NewRequest("POST", "", bytes.NewBuffer(jsonData))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	req.Header.Add("cookie", cookie)

	//send req to service
	serviceResp, err := forwardRequest(fmt.Sprintf("%v/cart/%v", os.Getenv("CART_SERVICE_URL"), productID), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	//read the response and write it
	resp, err := io.ReadAll(serviceResp.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Data(serviceResp.StatusCode, serviceResp.Header.Get("Content-Type"), resp)
}

// !
// UPDATE QUANTITY
func (x *CartHandler) UpdateQuantityOfProduct(c *gin.Context) {
	productID := c.Param("id")
	quantity := c.Query("quantity")
	//cookie
	cookie := c.Request.Header.Get("cookie")

	//
	var setExact CartQuantityRequest
	err := c.BindJSON(&setExact)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	jsonData, err := json.Marshal(&setExact)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	//
	req, err := http.NewRequest("PATCH", "", bytes.NewBuffer(jsonData))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	req.Header.Add("cookie", cookie)

	//
	serviceResp, err := forwardRequest(fmt.Sprintf("%v/cart/%v?quantity=%v", os.Getenv("CART_SERVICE_URL"), productID, quantity), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	data, err := io.ReadAll(serviceResp.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Data(serviceResp.StatusCode, serviceResp.Header.Get("Content-Type"), data)

}

// !
// DELETE PRODUCT ON CART
func (x *CartHandler) DeleteProductOnCart(c *gin.Context) {
	productID := c.Param("id")
	cookie := c.Request.Header.Get("cookie")

	//req
	req, err := http.NewRequest("DELETE", "", nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	req.Header.Add("cookie", cookie)
	//resp
	serviceResp, err := forwardRequest(fmt.Sprintf("%v/cart/%v", os.Getenv("CART_SERVICE_URL"), productID), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	data, err := io.ReadAll(serviceResp.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Data(serviceResp.StatusCode, serviceResp.Header.Get("Content-Type"), data)
}

// !
// CHECKOUT
func (x *CartHandler) CheckOut(c *gin.Context) {
	cookie := c.Request.Header.Get("cookie")

	//
	var addressID CartCheckOutType
	err := c.BindJSON(&addressID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	jsonData, err := json.Marshal(addressID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	//req
	req, err := http.NewRequest("POST", "", bytes.NewBuffer(jsonData))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	req.Header.Add("cookie", cookie)

	serviceResp, err := forwardRequest(fmt.Sprintf("%v/checkout", os.Getenv("CART_SERVICE_URL")), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	data, err := io.ReadAll(serviceResp.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Data(serviceResp.StatusCode, serviceResp.Header.Get("Content-Type"), data)
}
