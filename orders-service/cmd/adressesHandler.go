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
	userID, err := isCookieValid(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	//
	addresses, err := services.getAddresses(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": addresses,
	})

}

// !
// GET SINGLE ADDRESS
func (x *AdressesHandler) GetSingleAddressByID(c *gin.Context) {
	//get user id
	userID, err := isCookieValid(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	//get param
	addressIDStr := c.Param("id")

	//service
	address, err := services.getSingleAddress(userID, addressIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	//response
	c.JSON(http.StatusOK, gin.H{
		"data": address,
	})

}

// !
// ADD ADDRESS
func (x *AdressesHandler) AddNewAddress(c *gin.Context) {
	//get user info
	userID, err := isCookieValid(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	//get address infos as json
	var addressInfo NewAddress
	err = c.BindJSON(&addressInfo)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	addressInfo.UserID = userID

	//service
	err = services.addNewAddress(addressInfo)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	//response
	c.JSON(http.StatusCreated, gin.H{
		"message": "Address created succesfully",
	})

}

// !
// EDIT ADDRESS
func (x *AdressesHandler) EditAddressByID(c *gin.Context) {
	// get user info
	userID, err := isCookieValid(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	//get param
	addressIDStr := c.Param("id")

	// get address infos as json
	var addressInfo NewAddress
	err = c.BindJSON(&addressInfo)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
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
	userID, err := isCookieValid(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	//get param
	addressIDStr := c.Param("id")

	//service
	err = services.deleteAddressByID(addressIDStr, userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	//response
	c.JSON(http.StatusOK, gin.H{
		"message": "Address deleted succesfully",
	})

}
