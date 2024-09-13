package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()

	// init routes
	initRoutes(r)

	r.Run(":8081")
}

func connectDB() {

}
