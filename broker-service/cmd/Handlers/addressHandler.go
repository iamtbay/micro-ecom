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

type AddressHandler struct{}

func InitAddressHandlers() *AddressHandler {
	return &AddressHandler{}
}

// !
// GET ADDRESSES
func (x *AddressHandler) GetAddresses(c *gin.Context) {
	cookie := c.Request.Header.Get("cookie")

	req, err := http.NewRequest("GET", "", nil)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	req.Header.Add("cookie", cookie)

	serviceResp, err := forwardRequest(os.Getenv("ADRESSES_SERVICE_URL"), req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	defer serviceResp.Body.Close()
	resp, err := io.ReadAll(serviceResp.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.Data(serviceResp.StatusCode, serviceResp.Header.Get("Content-Type"), resp)
}

// !
// GET SINGLE ADDRESS
func (x *AddressHandler) GetSingleAddressByID(c *gin.Context) {
	id := c.Param("id")
	cookie := c.Request.Header.Get("cookie")

	req, err := http.NewRequest("GET", "", nil)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	req.Header.Add("cookie", cookie)

	serviceResp, err := forwardRequest(fmt.Sprintf("%v/%v", os.Getenv("ADRESSES_SERVICE_URL"), id), req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	defer serviceResp.Body.Close()
	resp, err := io.ReadAll(serviceResp.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.Data(serviceResp.StatusCode, serviceResp.Header.Get("Content-Type"), resp)
}

// !
// ADD NEW ADDRESS
func (x *AddressHandler) AddNewAddress(c *gin.Context) {
	cookie := c.Request.Header.Get("cookie")

	var newAddress NewAddress
	err := c.BindJSON(&newAddress)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	jsonData, err := json.Marshal(newAddress)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	req, err := http.NewRequest("POST", "", bytes.NewBuffer(jsonData))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	req.Header.Add("cookie", cookie)

	serviceResp, err := forwardRequest(os.Getenv("ADRESSES_SERVICE_URL"), req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	defer serviceResp.Body.Close()
	resp, err := io.ReadAll(serviceResp.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.Data(serviceResp.StatusCode, serviceResp.Header.Get("Content-Type"), resp)
}

// !
// EDIT ADDRESS
func (x *AddressHandler) EditAddressByID(c *gin.Context) {
	id := c.Param("id")
	cookie := c.Request.Header.Get("cookie")

	var newAddress NewAddress
	err := c.BindJSON(&newAddress)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	jsonData, err := json.Marshal(newAddress)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	req, err := http.NewRequest("PATCH", "", bytes.NewBuffer(jsonData))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	req.Header.Add("cookie", cookie)

	serviceResp, err := forwardRequest(fmt.Sprintf("%v/%v", os.Getenv("ADRESSES_SERVICE_URL"), id), req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	defer serviceResp.Body.Close()
	resp, err := io.ReadAll(serviceResp.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.Data(serviceResp.StatusCode, serviceResp.Header.Get("Content-Type"), resp)
}

// !
// DELETE ADDRESS BY ID
func (x *AddressHandler) DeleteAddressByID(c *gin.Context) {
	id := c.Param("id")
	cookie := c.Request.Header.Get("cookie")

	req, err := http.NewRequest("DELETE", "", nil)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	req.Header.Add("cookie", cookie)

	serviceResp, err := forwardRequest(fmt.Sprintf("%v/%v", os.Getenv("ADRESSES_SERVICE_URL"), id), req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	defer serviceResp.Body.Close()
	resp, err := io.ReadAll(serviceResp.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.Data(serviceResp.StatusCode, serviceResp.Header.Get("Content-Type"), resp)
}
