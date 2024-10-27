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

type ProductHandlers struct{}

func InitProductHandlers() *ProductHandlers {
	return &ProductHandlers{}
}

// get products
func (x *ProductHandlers) GetAllProducts(c *gin.Context) {
	//get page
	page := c.Query("page")
	if page == "" {
		page = "1"
	}

	//service
	serviceReq, err := http.NewRequest("GET", "", nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	//
	serviceResp, err := forwardRequest(fmt.Sprintf("%v?page=%v", os.Getenv("PRODUCT_SERVICE_URL"), page), serviceReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	defer serviceResp.Body.Close()
	//

	respBody, err := io.ReadAll(serviceResp.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.Data(http.StatusOK, serviceResp.Header.Get("Content-Type"), respBody)
}

// get single product
func (x *ProductHandlers) GetProductByID(c *gin.Context) {
	productID := c.Param("id")

	//request
	serviceResp, err := http.Get(fmt.Sprintf("%v/s/%v", os.Getenv("PRODUCT_SERVICE_URL"), productID))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	defer serviceResp.Body.Close()

	//
	resp, err := io.ReadAll(serviceResp.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.Data(serviceResp.StatusCode, serviceResp.Header.Get("Content-Type"), resp)

}

// !
// add product
func (x *ProductHandlers) AddProduct(c *gin.Context) {
	//get data
	var data ProductData
	err := c.BindJSON(&data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	//
	jsonData, err := json.Marshal(data)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	//get cookie
	cookie := c.Request.Header.Get("cookie")

	//create req
	req, err := http.NewRequest("POST", "", bytes.NewBuffer(jsonData))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	req.Header.Add("cookie", cookie)

	//service
	serviceResp, err := forwardRequest(fmt.Sprintf("%v/add", os.Getenv("PRODUCT_SERVICE_URL")), req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	defer serviceResp.Body.Close()

	//read response from service and send it to user
	resp, err := io.ReadAll(serviceResp.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.Data(serviceResp.StatusCode, serviceResp.Header.Get("Content-Type"), resp)
}

// !
// edit product
func (x *ProductHandlers) EditProduct(c *gin.Context) {
	var data ProductData
	productID := c.Param("id")
	if productID == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "unvalid id",
		})
		return
	}
	err := c.BindJSON(&data)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	jsonData, err := json.Marshal(data)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	//get cookie
	cookie := c.Request.Header.Get("cookie")

	//create req
	req, err := http.NewRequest("PATCH", "", bytes.NewBuffer(jsonData))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	req.Header.Add("cookie", cookie)
	//send to service
	serviceResp, err := forwardRequest(fmt.Sprintf("%v/s/%v", os.Getenv("PRODUCT_SERVICE_URL"), productID), req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	defer serviceResp.Body.Close()
	//read response and send to the user
	resp, err := io.ReadAll(serviceResp.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.Data(serviceResp.StatusCode, serviceResp.Header.Get("Content-Type"), resp)

}

// !
// delete product
func (x *ProductHandlers) DeleteProduct(c *gin.Context) {
	productID := c.Param("id")
	if productID == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "unvalid id",
		})
		return
	}

	//get cookie
	cookie := c.Request.Header.Get("cookie")

	//create req
	req, err := http.NewRequest("DELETE", "", nil)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	req.Header.Add("cookie", cookie)
	//send to service
	serviceResp, err := forwardRequest(fmt.Sprintf("%v/s/%v", os.Getenv("PRODUCT_SERVICE_URL"), productID), req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	defer serviceResp.Body.Close()
	//read response and send to the user
	resp, err := io.ReadAll(serviceResp.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.Data(serviceResp.StatusCode, serviceResp.Header.Get("Content-Type"), resp)
}
