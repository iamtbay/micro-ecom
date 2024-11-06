package main

import (
	"log"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/gin-gonic/gin"
	"github.com/rabbitmq/amqp091-go"
)

var (
	es *elasticsearch.Client
	ch *amqp091.Channel
)

func main() {

	r := gin.Default()
	var err error
	es, err = elasticsearch.NewDefaultClient()

	if err != nil {
		log.Fatalf("Error creating elastic client %s", err)
	}

	initRoutes(r)

	connectRabbitMQ()
	consumeMessages()
	//
	r.Run(":8087")

}
