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

type InventoryHandler struct{}

func InitInventoryHandlers() *InventoryHandler {
	return &InventoryHandler{}
}

// GET PRODUCT STOCK
func (x *InventoryHandler) GetProductStock(c *gin.Context) {
	cookie := c.Request.Header.Get("cookie")
	productID := c.Param("id")

	req, err := http.NewRequest("GET", "", nil)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	req.Header.Add("cookie", cookie)

	serviceResp, err := forwardRequest(fmt.Sprintf("%v/%v", os.Getenv("INVENTORY_SERVICE_URL"), productID), req)
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
// ADD NEW PRODUCT STOCK
func (x *InventoryHandler) AddNewProductStock(c *gin.Context) {
	cookie := c.Request.Header.Get("cookie")
	var productInfo InventoryProduct
	err := c.BindJSON(&productInfo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	jsonData, _ := json.Marshal(productInfo)

	req, _ := http.NewRequest("POST", "", bytes.NewBuffer(jsonData))

	req.Header.Add("cookie", cookie)
	serviceResp, err := forwardRequest(fmt.Sprintf("%v/new", os.Getenv("INVENTORY_SERVICE_URL")), req)
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
// RE STOCK PRODUCT
func (x *InventoryHandler) RestockProduct(c *gin.Context) {
	cookie := c.Request.Header.Get("cookie")
	productID := c.Param("id")
	var productInfo InventoryProduct
	err := c.BindJSON(&productInfo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	jsonData, _ := json.Marshal(productInfo)

	req, _ := http.NewRequest("PATCH", "", bytes.NewBuffer(jsonData))

	req.Header.Add("cookie", cookie)
	serviceResp, err := forwardRequest(fmt.Sprintf("%v/restock/%v", os.Getenv("INVENTORY_SERVICE_URL"), productID), req)
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
// CANCEL STOCK
func (x *InventoryHandler) CancelStockReservation(c *gin.Context) {
	cookie := c.Request.Header.Get("cookie")

	var productInfo InventoryProductSale
	err := c.BindJSON(&productInfo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	productInfo.ProductID = c.Param("id")
	jsonData, _ := json.Marshal(productInfo)

	req, _ := http.NewRequest("PATCH", "", bytes.NewBuffer(jsonData))

	req.Header.Add("cookie", cookie)
	serviceResp, err := forwardRequest(fmt.Sprintf("%v/cancel/%v", os.Getenv("INVENTORY_SERVICE_URL"), productInfo.ProductID), req)
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
// CONFIRM STOCK RESERVATION
func (x *InventoryHandler) ConfirmStockReservation(c *gin.Context) {
	cookie := c.Request.Header.Get("cookie")
	var productInfo InventoryProductSale
	err := c.BindJSON(&productInfo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	productInfo.ProductID = c.Param("id")
	jsonData, _ := json.Marshal(productInfo)

	req, _ := http.NewRequest("PATCH", "", bytes.NewBuffer(jsonData))

	req.Header.Add("cookie", cookie)
	serviceResp, err := forwardRequest(fmt.Sprintf("%v/reserved/%v", os.Getenv("INVENTORY_SERVICE_URL"), productInfo.ProductID), req)
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
// UPDATE STOCK AFTER SALE
func (x *InventoryHandler) UpdateStockAfterSale(c *gin.Context) {
	cookie := c.Request.Header.Get("cookie")
	productID := c.Param("id")
	var productInfo InventoryProductSale
	err := c.BindJSON(&productInfo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	productInfo.ProductID = productID
	jsonData, _ := json.Marshal(productInfo)

	req, _ := http.NewRequest("PATCH", "", bytes.NewBuffer(jsonData))

	req.Header.Add("cookie", cookie)
	serviceResp, err := forwardRequest(fmt.Sprintf("%v/sold/%v", os.Getenv("INVENTORY_SERVICE_URL"), productID), req)
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
