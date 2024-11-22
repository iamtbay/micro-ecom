package main

import (
	"context"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	mongoConn  *mongo.Client
	collection *mongo.Collection
)

func main() {
	r := gin.Default()
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	connectDB()
	initRoutes(r)

	err := r.Run(":8087")
	if err != nil {
		panic(err)
	}
}

func connectDB() {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(os.Getenv("MONGO_URI")).SetServerAPIOptions(serverAPI)

	var err error
	mongoConn, err = mongo.Connect(context.Background(), opts)
	if err != nil {
		panic(err)
	}

	var result bson.M

	if err := mongoConn.Database("go-fs-ecommerce").RunCommand(context.Background(), bson.D{{
		"ping", 1}}).Decode(&result); err != nil {
		panic(err)
	}

	fmt.Println("DB PINGED AND SUCCESSFULLY CONNECTED")
	collection = mongoConn.Database("go-fs-ecommerce").Collection("favorites")

}
