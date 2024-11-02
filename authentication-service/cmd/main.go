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
		log.Fatal("error while loading .env file")
	}
	//connect db
	connectDB()
	defer conn.Close()

	//init routes
	initRoutes(r)

	//RUN SERVER
	r.Run(":8081")
}

// CONNECT DB
func connectDB() {
	var err error
	conn, err = pgxpool.New(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connection to database %v \n", err)
		os.Exit(1)
	}
	fmt.Println("Database Ready!!!")

}
