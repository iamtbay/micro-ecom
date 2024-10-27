package main

import (
	"context"
	"fmt"
	"log"
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
		log.Fatalf("Somethin went wrong while .env file mounting %v")
		os.Exit(1)
	}
	
	initRoutes(r)
	err := connectDB()
	if err != nil {
		panic(err)
	}

	r.Run(":8085")
}

func connectDB() error {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(os.Getenv("MONGO_URI")).SetServerAPIOptions(serverAPI)
	var err error
	mongoConn, err = mongo.Connect(context.Background(), opts)
	if err != nil {
		return err
	}
	var result bson.M
	if err := mongoConn.Database("go-fs-ecommerce").RunCommand(context.TODO(), bson.D{{"ping", 1}}).Decode(&result); err != nil {
		panic(err)
	}
	fmt.Println("DB pinged succesfully and connected Mongo!")
	collection = mongoConn.Database("go-fs-ecommerce").Collection("reviews")
	return nil
}
