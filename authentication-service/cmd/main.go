package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

var conn *pgx.Conn

func main() {
	r := gin.Default()
	if err := godotenv.Load(); err != nil {
		fmt.Println(err)
		log.Fatal("error loading .env file")
	}
	//connect db
	connectDB()
	defer conn.Close(context.Background())
	//init routes
	initRoutes(r)

	//RUN SERVER
	r.Run()
}

// CONNECT DB
func connectDB() {
	var err error
	conn, err = pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connection to database %v \n", err)
		os.Exit(1)
	}
	fmt.Println("Database Ready!!!")

}
