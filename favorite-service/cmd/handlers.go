package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handlers struct{}

func initHandlers() *Handlers {
	return &Handlers{}
}

var services = initService()

func (x *Handlers) newFavorite(c *gin.Context) {
	userID, _ := isCookieValid(c)
	productID := c.Param("id")

	//
	err := services.newFavorite(userID, productID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": "Product added favorites",
	})

}

func (x *Handlers) getAllFavorites(c *gin.Context) {
	userID, _ := isCookieValid(c)
	favorites, err := services.getAllFavorites(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": favorites,
	})
}

func (x *Handlers) removeFavorite(c *gin.Context) {
	userID, _ := isCookieValid(c)

	productID := c.Param("id")
	err := services.removeFavorite(userID, productID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": "Product removed your favorite list.",
	})
}
