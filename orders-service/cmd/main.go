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

var conn *pgxpool.Pool

func main() {
	r := gin.Default()

	if err := godotenv.Load(); err != nil {
		log.Fatal("error loading .env file")
	}

	// init routes
	initRoutes(r)

	//connect to db
	connectDB()

	//rabbitmq
	ch, err := connectRabbitMQ()
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}
	defer ch.Close()
	consumeMessages(ch)

	r.Run(":8083")
}

func connectDB() {
	var err error
	conn, err = pgxpool.New(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect database %v \n", err)
		os.Exit(1)
	}
}
