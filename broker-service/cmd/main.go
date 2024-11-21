package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	r := gin.Default()

	//godotenv
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	initRoutes(r)

	//start
	r.Run(":8080")
}
