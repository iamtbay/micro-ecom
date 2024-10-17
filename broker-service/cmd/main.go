package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"log"
)

func main() {
	r := gin.Default()

	//godotenv
	if err := godotenv.Load(); err != nil {
		log.Fatal(".env file couldn't find !!!")
	}

	initRoutes(r)

	//start tv
	r.Run()
}
