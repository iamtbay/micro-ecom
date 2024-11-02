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
	//GODOT ENV
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	//CONNECT DB
	err := connectDB()
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := mongoConn.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	//ROUTES
	initRoutes(r)
	connectRabbitMQ()
	defer ch.Close()

	//SERVER
	err = r.Run(":8082")
	if err != nil {
		panic(err)
	}

}

// connect db
func connectDB() error {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(os.Getenv("MONGO_URI")).SetServerAPIOptions(serverAPI)

	var err error
	mongoConn, err = mongo.Connect(context.TODO(), opts)
	if err != nil {
		return err
	}

	var result bson.M

	if err := mongoConn.Database("go-fs-ecommerce").RunCommand(context.TODO(), bson.D{{"ping", 1}}).Decode(&result); err != nil {
		panic(err)
	}
	fmt.Println("DB pinged succesfully and connected Mongo!")

	collection = mongoConn.Database("go-fs-ecommerce").Collection("products")

	return nil
}
