package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handlers struct{}

var services = initServices()

func initHandlers() *Handlers {
	return &Handlers{}
}

// CHECK
func (x *Handlers) check(c *gin.Context) {
	cookie, _ := getCookie(c)

	userInfo, err := services.checkUser(cookie)
	if err != nil {
		c.JSON(401, gin.H{
			"error":   err.Error(),
			"message": "Your session expired, please log in again.",
		})
		return
	}

	//
	c.JSON(200, gin.H{
		"message": "User has verified",
		"data": gin.H{
			"user_id": userInfo.ID,
			"name":    userInfo.Name,
			"surname": userInfo.Surname,
			"email":   userInfo.Email,
		},
	})
}

// !
// LOGIN
func (x *Handlers) login(c *gin.Context) {
	var userInfo UserBasicInfo
	err := c.BindJSON(&userInfo)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"message": "Error while reading user informations",
		})
		return
	}

	//service req
	token, userInfos, err := services.login(userInfo)
	if err != nil {
		c.JSON(401, gin.H{
			"message": "Something went wrong while logging in.",
			"error":   err.Error(),
		})
		return
	}
	//arrange jwt as a cookie
	setCookie(c, "accessToken", token)

	//RETURN RESPONSE
	c.JSON(http.StatusOK, gin.H{
		"message": "User has verified",
		"data": gin.H{
			"user_id": userInfos.ID,
			"name":    userInfos.Name,
			"surname": userInfos.Surname,
			"email":   userInfos.Email,
		},
	})

}

// !
// SIGNUP
func (x *Handlers) signup(c *gin.Context) {
	var userInfos UserBasicInfo
	err := c.BindJSON(&userInfos)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"message": "Error while reading user informations",
		})
		return
	}

	//service req
	err = services.signup(userInfos)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"message": "User couldn't signed up",
		})
		return
	}
	//RETURN RESPONSE
	c.JSON(http.StatusCreated, gin.H{
		"message": "User has succesfully registered",
		"data": gin.H{
			"name":    userInfos.Name,
			"surname": userInfos.Surname,
			"email":   userInfos.Email,
		},
	})
}

// !
// EDIT
func (x *Handlers) edit(c *gin.Context) {
	var userInfos UserBasicInfo
	err := c.BindJSON(&userInfos)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"message": "Error while reading user informations",
		})
		return
	}

	//get token
	token, err := getCookie(c)
	if err != nil {
		c.JSON(401, gin.H{
			"error": err.Error(),
		})
		return
	}

	//service req
	err = services.edit(userInfos, token)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	//RETURN RESPONSE
	c.JSON(http.StatusOK, gin.H{
		"message": "User infos are updated",
		"data": gin.H{
			"name":    userInfos.Name,
			"surname": userInfos.Surname,
			"email":   userInfos.Email,
		},
	})
}

// !
// CHANGE PASSWORD
func (x *Handlers) changePassword(c *gin.Context) {
	token, _ := getCookie(c)

	var newPassword NewPassword
	err := c.BindJSON(&newPassword)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"message": "Error while reading user informations",
		})
		return
	}

	//service req
	err = services.changePassword(newPassword.Password, token)
	if err != nil {
		c.JSON(400, gin.H{
			"error":   err.Error(),
			"message": "Error while changing password",
		})
		return
	}
	//RETURN RESPONSE
	c.JSON(http.StatusNoContent, gin.H{
		"message": "User password changed",
		"data":    nil,
	})
}

// !
// DELETE
func (x *Handlers) delete(c *gin.Context) {
	//get cookie
	token, _ := getCookie(c)

	//service
	err := services.delete(token)
	if err != nil {
		c.JSON(401, gin.H{
			"error": err.Error(),
		})
		return
	}

	//return json
	c.JSON(http.StatusNoContent, gin.H{
		"message": "User has deleted succesfully",
	})
}

// !
// LOGOUT
func (x *Handlers) logout(c *gin.Context) {
	c.SetCookie("accessToken", "", -1, "/", "localhost", true, true)
	c.JSON(http.StatusNoContent, gin.H{
		"message": "User succesfully logged out.",
	})
}
