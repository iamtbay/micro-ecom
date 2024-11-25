package main

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/chai2010/webp"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
	fmt.Println("Prod", productInfo)

	//service
	product, err := services.addProduct(&productInfo)
	if err != nil {
		errJSON(http.StatusInternalServerError, err, "Something went wrong while adding product", c)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Product succesfully added",
		"data":    product,
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

//!
// ADD IMAGES

func (x *ProductsHandler) addImages(c *gin.Context) {
	productIdStr := c.Param("id")
	productID, err := primitive.ObjectIDFromHex(productIdStr)
	if err != nil {
		errJSON(400, err, "error while reading objectID", c)
		return
	}
	var images []string
	form, err := c.MultipartForm()
	if err != nil {
		errJSON(400, err, "error getting file", c)
		return
	}

	files := form.File["images"]

	for _, file := range files {
		uniqueName := uuid.New()
		srcFile, err := file.Open()
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		defer srcFile.Close()

		ext := strings.ToLower(filepath.Ext(file.Filename))

		var img image.Image
		switch ext {
		case ".webp":
			webpPath := filepath.Join("./images/productImages/", fmt.Sprintf("%v.webp", uniqueName))
			if err := c.SaveUploadedFile(file, webpPath); err != nil {
				c.JSON(400, gin.H{"error": err.Error()})
				return
			}
			images = append(images, fmt.Sprintf("%v.webp", uniqueName))
			continue
		case ".jpg", ".jpeg":
			img, err = jpeg.Decode(srcFile)
			if err != nil {
				c.JSON(400, gin.H{"error": err.Error()})
				return
			}
		case ".png":
			img, err = png.Decode(srcFile)
			if err != nil {
				c.JSON(400, gin.H{"error": err.Error()})
				return
			}

		default:
			c.String(http.StatusBadRequest, "unsupported image format")
		}
		err = x.saveImage(uniqueName, img, c)
		images = append(images, fmt.Sprintf("%v.webp", uniqueName))
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
	}

	//productId?
	services.addImages(images, productID)
	c.JSON(http.StatusOK, gin.H{
		"message": "images uploaded",
		"images":  images,
	})

}

// save img method
func (x *ProductsHandler) saveImage(uniqueName uuid.UUID, img image.Image, c *gin.Context) error {
	webpPath := filepath.Join("./images/productImages", fmt.Sprintf("%v.webp", uniqueName))
	outFile, err := os.Create(webpPath)
	if err != nil {
		return err
	}

	defer outFile.Close()

	if err := webp.Encode(outFile, img, &webp.Options{Lossless: false, Quality: 80}); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return err
	}
	return nil
}

// RETURN ERROR JSON
func errJSON(status int, err error, message string, c *gin.Context) {
	c.JSON(status, gin.H{
		"message": message,
		"error":   err.Error(),
	})

}
