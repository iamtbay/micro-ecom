package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/joho/godotenv"
	"github.com/rabbitmq/amqp091-go"
)

// CART SERVICE
var (
	rdb *redis.Client
	ch  *amqp091.Channel
)

func main() {
	r := gin.Default()
	//.env
	if err := godotenv.Load(); err != nil {
		log.Fatal("error loading .env file")
	}

	connectDB()
	initRoutes(r)

	var err error
	//connect rabbitmq
	ch, err = connectRabbitMQ()
	if err != nil {
		log.Fatalf("Failed to connect to rabbitmq %v", err)
	}
	defer ch.Close()

	consumeMessages()

	r.Run(":8084")
}

// CONNECT DB
func connectDB() {
	rdb = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
}
