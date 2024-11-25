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

type AuthHandlers struct{}

func InitAuthHandlers() *AuthHandlers {
	return &AuthHandlers{}
}

// !
// Check
func (x *AuthHandlers) Check(c *gin.Context) {
	cookie := c.Request.Header.Get("cookie")

	req, err := http.NewRequest("GET", "", nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating request"})
		return
	}
	req.Header.Add("cookie", cookie)

	serviceResp, err := forwardRequest(fmt.Sprintf("%v/check", os.Getenv("AUTH_SERVICE_URL")), req)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	defer serviceResp.Body.Close()

	respBody, err := io.ReadAll(serviceResp.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read response from service"})
		return
	}

	c.Data(serviceResp.StatusCode, serviceResp.Header.Get("Content-Type"), respBody)
}

// !
// LOGIN
func (x *AuthHandlers) Login(c *gin.Context) {
	var data AuthRequest
	err := c.BindJSON(&data)
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return
	}
	//marshal
	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return
	}
	req, _ := http.NewRequest("POST", "", bytes.NewReader(jsonData))
	serviceResp, err := forwardRequest(fmt.Sprintf("%v/login", os.Getenv("AUTH_SERVICE_URL")), req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer serviceResp.Body.Close()

	// read the body from service
	respBody, err := io.ReadAll(serviceResp.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read response from service"})
		return
	}

	//cookie
	if cookies := serviceResp.Header["Set-Cookie"]; len(cookies) > 0 {
		for _, cookie := range cookies {
			c.Writer.Header().Add("Set-Cookie", cookie)
		}
	}

	//send response
	c.Data(serviceResp.StatusCode, serviceResp.Header.Get("Content-Type"), respBody)
}

// !
// Signup
func (x *AuthHandlers) Signup(c *gin.Context) {
	// turn json
	var data AuthRequest
	err := c.BindJSON(&data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// make new req
	req, err := http.NewRequest("POST", "", bytes.NewReader(jsonData))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to request"})
		return
	}

	// send req
	serviceResp, err := forwardRequest(fmt.Sprintf("%v/signup", os.Getenv("AUTH_SERVICE_URL")), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer serviceResp.Body.Close()

	// read the response
	respBody, err := io.ReadAll(serviceResp.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read response from service"})
		return
	}

	c.Data(serviceResp.StatusCode, serviceResp.Header.Get("Content-Type"), respBody)
}

// !
// EDIT
func (x *AuthHandlers) Edit(c *gin.Context) {

	// turn json
	var data AuthRequest
	err := c.BindJSON(&data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	//turn it to json
	jsonData, err := json.Marshal(data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to turning json"})
		return
	}

	//get cookie
	cookie := c.Request.Header.Get("cookie")

	//make req
	req, err := http.NewRequest("PATCH", "", bytes.NewBuffer(jsonData))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	req.Header.Add("cookie", cookie)

	//send req
	serviceResp, err := forwardRequest(fmt.Sprintf("%v/edit", os.Getenv("AUTH_SERVICE_URL")), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	defer serviceResp.Body.Close()

	//read the response
	respBody, err := io.ReadAll(serviceResp.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read response from service"})
		return
	}
	c.Data(serviceResp.StatusCode, serviceResp.Header.Get("Content-Type"), respBody)
}

// !
// CHANGE-PASSWORD
func (x *AuthHandlers) ChangePassword(c *gin.Context) {
	// turn json
	var data AuthChangePassword
	err := c.BindJSON(&data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// turn to json
	jsonData, err := json.Marshal(data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to return JSON"})
		return
	}
	// get cookie
	cookie := c.Request.Header.Get("cookie")

	// set new req and add header cookie
	req, err := http.NewRequest("PATCH", "", bytes.NewBuffer(jsonData))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to creting request"})
		return
	}
	req.Header.Add("cookie", cookie)
	// send to service
	serviceResp, err := forwardRequest(fmt.Sprintf("%v/changepassword", os.Getenv("AUTH_SERVICE_URL")), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer serviceResp.Body.Close()
	// read the service response
	resp, err := io.ReadAll(serviceResp.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to reading resp"})
		return
	}
	// forward the response
	c.Data(serviceResp.StatusCode, serviceResp.Header.Get("Content-Type"), resp)
}

// !
// DELETE
func (x *AuthHandlers) Delete(c *gin.Context) {
}

// !
// LOGOUT
func (x *AuthHandlers) Logout(c *gin.Context) {
	cookie := c.Request.Header.Get("cookie")

	req, err := http.NewRequest("POST", "", nil)
	if err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}
	req.Header.Add("cookie", cookie)

	serviceResp, err := forwardRequest(fmt.Sprintf("%v/logout", os.Getenv("AUTH_SERVICE_URL")), req)
	if err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}
	defer serviceResp.Body.Close()

	respBody, err := io.ReadAll(serviceResp.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read response from service"})
		return
	}

	//cookie
	if cookies := serviceResp.Header["Set-Cookie"]; len(cookies) > 0 {
		for _, cookie := range cookies {
			c.Writer.Header().Add("Set-Cookie", cookie)
		}
	}

	c.Data(serviceResp.StatusCode, serviceResp.Header.Get("Content-Type"), respBody)
}
