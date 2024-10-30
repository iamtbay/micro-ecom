package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

var (
	conn *pgxpool.Pool
)

func main() {
	r := gin.Default()
	if err := godotenv.Load(); err != nil {
		log.Panicf("error while .env file loading, %v \n", err)
		return
	}

	connectDB()
	connectRabbitMQ()

	initRoutes(r)
	r.Run(":8086")
}

// CONNECT DB
func connectDB() {
	var err error
	conn, err = pgxpool.New(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect Database %v \n", err)
		os.Exit(1)
	}
}
