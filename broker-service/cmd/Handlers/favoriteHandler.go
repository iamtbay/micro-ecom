package handlersPackage

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type FavoriteHandler struct{}

func InitFavoriteHandler() *FavoriteHandler {
	return &FavoriteHandler{}
}

func (x *FavoriteHandler) AddToFavoriteList(c *gin.Context) {
	cookie := c.Request.Header.Get("cookie")
	productID := c.Param("id")
	req, err := http.NewRequest("POST", "", nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	req.Header.Add("cookie", cookie)

	serviceResp, err := forwardRequest(fmt.Sprintf("%v/favorites/%v", os.Getenv("FAVORITE_SERVICE_URL"), productID), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	resp, err := io.ReadAll(serviceResp.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Data(serviceResp.StatusCode, serviceResp.Header.Get("Content-Type"), resp)
}

//REMOVE FROM FAV LIST
func (x *FavoriteHandler) RemoveFromFavoriteList(c *gin.Context) {
	cookie := c.Request.Header.Get("cookie")
	productID := c.Param("id")

	req, err := http.NewRequest("DELETE", "", nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	req.Header.Add("cookie", cookie)

	serviceResp, err := forwardRequest(fmt.Sprintf("%v/favorites/%v", os.Getenv("FAVORITE_SERVICE_URL"), productID), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	resp, err := io.ReadAll(serviceResp.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Data(serviceResp.StatusCode, serviceResp.Header.Get("Content-Type"), resp)
}

//GET FAVORITE LIST
func (x *FavoriteHandler) GetFavoriteList(c *gin.Context) {
	cookie := c.Request.Header.Get("cookie")

	req, err := http.NewRequest("GET", "", nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	req.Header.Add("cookie", cookie)

	serviceResp, err := forwardRequest(fmt.Sprintf("%v/favorites", os.Getenv("FAVORITE_SERVICE_URL")), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	resp, err := io.ReadAll(serviceResp.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Data(serviceResp.StatusCode, serviceResp.Header.Get("Content-Type"), resp)
}
