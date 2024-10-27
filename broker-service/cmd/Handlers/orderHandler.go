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

type OrderHandler struct{}

func InitOrderHandlers() *OrderHandler {
	return &OrderHandler{}
}

// !
// GET ORDERS BY USER ID
func (x *OrderHandler) GetOrdersByUserID(c *gin.Context) {
	id := c.Param("id")
	//get cookie
	cookie := c.Request.Header.Get("cookie")

	//
	req, err := http.NewRequest("GET", "", nil)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	req.Header.Add("cookie", cookie)

	//
	serviceResp, err := forwardRequest(fmt.Sprintf("%v/user/%v", os.Getenv("ORDERS_SERVICE_URL"), id), req)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	defer serviceResp.Body.Close()

	//
	resp, err := io.ReadAll(serviceResp.Body)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	c.Data(serviceResp.StatusCode, serviceResp.Header.Get("Content-Type"), resp)
}

// !
// GET ORDER BY ID
func (x *OrderHandler) GetSingleOrder(c *gin.Context) {
	//get cookie
	id := c.Param("id")
	cookie := c.Request.Header.Get("cookie")

	//
	req, err := http.NewRequest("GET", "", nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	req.Header.Add("cookie", cookie)

	//
	serviceResp, err := forwardRequest(fmt.Sprintf("%v/%v", os.Getenv("ORDERS_SERVICE_URL"), id), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer serviceResp.Body.Close()

	//
	resp, err := io.ReadAll(serviceResp.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Data(serviceResp.StatusCode, serviceResp.Header.Get("Content-Type"), resp)
}

// !
// NEW ORDER
func (x *OrderHandler) NewOrder(c *gin.Context) {
	cookie := c.Request.Header.Get("cookie")

	var newOrder OrderRequest
	err := c.BindJSON(&newOrder)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	jsonData, err := json.Marshal(newOrder)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	req, err := http.NewRequest("POST", "", bytes.NewBuffer(jsonData))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	req.Header.Add("cookie", cookie)

	serviceResp, err := forwardRequest(fmt.Sprintf("%v/newOrder", os.Getenv("ORDERS_SERVICE_URL")), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer serviceResp.Body.Close()

	resp, err := io.ReadAll(serviceResp.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Data(serviceResp.StatusCode, serviceResp.Header.Get("Content-Type"), resp)
}

// !
// DELETE ORDER
func (x *OrderHandler) DeleteOrder(c *gin.Context) {
	//get cookie
	id := c.Param("id")
	cookie := c.Request.Header.Get("cookie")

	//
	req, err := http.NewRequest("DELETE", "", nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	req.Header.Add("cookie", cookie)

	//
	serviceResp, err := forwardRequest(fmt.Sprintf("%v/%v", os.Getenv("ORDERS_SERVICE_URL"), id), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer serviceResp.Body.Close()

	//
	resp, err := io.ReadAll(serviceResp.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Data(serviceResp.StatusCode, serviceResp.Header.Get("Content-Type"), resp)
}
