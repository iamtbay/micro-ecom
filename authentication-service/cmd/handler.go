package main

import (
	"fmt"
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
	fmt.Println("check")
}

// LOGIN
func (x *Handlers) login(c *gin.Context) {
	var userInfos UserBasicInfo
	err := c.BindJSON(&userInfos)
	if err != nil {
		fmt.Println("Error")
		return
	}

	//service req
	token, err := services.login(userInfos)
	if err != nil {
		c.JSON(401, gin.H{
			"error": err.Error(),
		})
		return
	}
	//arrange jwt as a cookie
	setCookie(c, "accessToken", token)

	//RETURN RESPONSE
	successJSON(http.StatusOK, "succesfully logged in", userInfos.Name, userInfos.Surname, userInfos.Email, c)

}

// SIGNUP
func (x *Handlers) signup(c *gin.Context) {
	var userInfos UserBasicInfo
	err := c.BindJSON(&userInfos)
	if err != nil {
		panic(err)

	}

	//service req
	err = services.signup(userInfos)
	if err != nil {
		c.JSON(401, gin.H{
			"error": err.Error(),
		})
		return
	}
	successJSON(http.StatusCreated, "succesfully registered", userInfos.Name, userInfos.Surname, userInfos.Email, c)
}

// EDIT
func (x *Handlers) edit(c *gin.Context) {
	fmt.Println("edit")
	var userInfos UserBasicInfo
	err := c.BindJSON(&userInfos)
	if err != nil {
		fmt.Println("Error")
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
		c.JSON(401, gin.H{
			"error": err.Error(),
		})
		return
	}
	successJSON(http.StatusAccepted, "user infos updated", userInfos.Name, userInfos.Surname, userInfos.Email, c)
}

// CHANGE PASSWORD
func (x *Handlers) changePassword(c *gin.Context) {
	token, err := getCookie(c)
	if err != nil {
		c.JSON(401, gin.H{
			"error": err.Error(),
		})
		return
	}

	var newPassword NewPassword
	err = c.BindJSON(&newPassword)
	if err != nil {
		fmt.Println("Error")
		return
	}

	//service req
	err = services.changePassword(newPassword.Password, token)
	if err != nil {
		c.JSON(401, gin.H{
			"error": err.Error(),
		})
		return
	}
	successJSON(http.StatusAccepted, "user password changed", "", "", "", c)
}

// DELETE
func (x *Handlers) delete(c *gin.Context) {
	fmt.Println("delete")
	//get cookie
	token, err := getCookie(c)
	if err != nil {
		c.JSON(401, gin.H{
			"error": err.Error(),
		})
		return
	}
	//service
	err = services.delete(token)
	if err != nil {
		c.JSON(401, gin.H{
			"error": err.Error(),
		})
		return
	}

	//return json
	c.JSON(http.StatusOK, gin.H{
		"message": "user deactiveted!",
	})
}

func (x *Handlers) logout(c *gin.Context) {

	c.SetCookie("accessToken", "", -1, "/", "localhost", true, true)
	c.JSON(http.StatusOK, gin.H{
		"message": "Succesfully logged out.",
	})
}

// return err json
func successJSON(code int, message, name, surname, email string, c *gin.Context) {
	c.JSON(code, gin.H{
		"message": message,
		"name":    name,
		"surname": surname,
		"email":   email,
	})
}
