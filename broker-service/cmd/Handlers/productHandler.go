package handlersPackage

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"

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
// add images
func (x *ProductHandlers) AddImages(c *gin.Context) {
	//
	form, err := c.MultipartForm()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	files := form.File["images"]
	x.sendFilesToProductService(files, c)

}

func (x *ProductHandlers) sendFilesToProductService(files []*multipart.FileHeader, c *gin.Context) error {
	productID := c.Param("id")
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	for _, file := range files {
		srcFile, err := file.Open()
		if err != nil {
			return fmt.Errorf("failed to open file %v", err)
		}
		defer srcFile.Close()

		//
		part, err := writer.CreateFormFile("images", filepath.Base(file.Filename))
		if err != nil {
			return err
		}

		_, err = io.Copy(part, srcFile)
		if err != nil {
			return err
		}

	}

	err := writer.Close()
	if err != nil {
		return err
	}
	req, err := http.NewRequest("POST", fmt.Sprintf("%v/image/add/%v", os.Getenv("PRODUCT_SERVICE_URL"), productID), body)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())
	req.Header.Add("cookie", c.Request.Header.Get("cookie"))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	respI, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	c.Data(resp.StatusCode, resp.Header.Get("Content-Type"), respI)
	fmt.Println("Images uploaded!")
	return nil
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
