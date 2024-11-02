package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type AdressesHandler struct{}

func initAdressesHandler() *AdressesHandler {
	return &AdressesHandler{}
}

// !
// GET ADDRESSES
func (x *AdressesHandler) GetAddresses(c *gin.Context) {
	//get user id
	userID, _ := isCookieValid(c)

	//
	addresses, err := services.getAddresses(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"message": "Error while getting adresses",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data":    addresses,
		"message": "Adresses found",
	})

}

// !
// GET SINGLE ADDRESS
func (x *AdressesHandler) GetSingleAddressByID(c *gin.Context) {
	//get user id
	userID, _ := isCookieValid(c)

	//get param
	addressIDStr := c.Param("id")

	//service
	address, err := services.getSingleAddress(userID, addressIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"message": "Adress couldn't find",
		})
		return
	}

	//response
	c.JSON(http.StatusOK, gin.H{
		"data":    address,
		"message": "Adress found",
	})

}

// !
// ADD ADDRESS
func (x *AdressesHandler) AddNewAddress(c *gin.Context) {
	//get user info
	userID, _ := isCookieValid(c)

	//get address infos as json
	var addressInfo NewAddress
	err := c.BindJSON(&addressInfo)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"message": "Error while reading adress informations",
		})
		return
	}
	addressInfo.UserID = userID

	//service
	err = services.addNewAddress(addressInfo)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"message": "Error while adding new adress",
		})
		return
	}

	//response
	c.JSON(http.StatusOK, gin.H{
		"message": "Address created succesfully",
	})

}

// !
// EDIT ADDRESS
func (x *AdressesHandler) EditAddressByID(c *gin.Context) {
	// get user info
	userID, _ := isCookieValid(c)

	//get param
	addressIDStr := c.Param("id")

	// get address infos as json
	var addressInfo NewAddress
	err := c.BindJSON(&addressInfo)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"message": "Error while reading adress informations",
		})
		return
	}
	addressInfo.UserID = userID

	//service
	err = services.editAddressByID(addressInfo, addressIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	//response
	c.JSON(http.StatusOK, gin.H{
		"message": "Address succesfully updated",
	})
}

// !
// DELETE ADDRESS
func (x *AdressesHandler) DeleteAddressByID(c *gin.Context) {
	// get user info
	userID, _ := isCookieValid(c)

	//get param
	addressIDStr := c.Param("id")

	//service
	err := services.deleteAddressByID(addressIDStr, userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"message": "Error while deleting an adress",
		})
		return
	}

	//response
	c.JSON(http.StatusOK, gin.H{
		"message": "Address deleted succesfully",
	})

}
