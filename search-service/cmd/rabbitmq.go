package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/rabbitmq/amqp091-go"
)

func connectRabbitMQ() {
	conn, err := amqp091.Dial("amqp://guest:guest@localhost:5672/")
	failOnErr(err, "Connection error")
	//

	ch, err = conn.Channel()
	failOnErr(err, "Channel error")

	err = ch.ExchangeDeclare(
		"search_exchange",
		"direct",
		true,
		false,
		false,
		false,
		nil,
	)
	failOnErr(err, "Exchange declare")

	q, err := ch.QueueDeclare(
		"search_queue",
		true,
		false,
		false,
		false,
		nil,
	)
	failOnErr(err, "Queue Declare")

	err = ch.QueueBind(
		q.Name,
		"index.new",
		"search_exchange",
		false,
		nil,
	)
	failOnErr(err, "Queue Bind")
}

func consumeMessages() {
	msgs, err := ch.Consume(
		"search_queue",
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	failOnErr(err, "Consume error")

	go func() {
		for msg := range msgs {
			var product Product
			err := json.Unmarshal(msg.Body, &product)
			if err != nil {
				fmt.Println("Error while unmarshal", err)
				return
			}
			err = services.IndexProduct(product)
			if err!=nil{
				fmt.Println("Error while indexing", err)
				return
			}
		}
	}()
}

func failOnErr(err error, msg string) {
	if err != nil {
		log.Fatalf("%v: %v", msg, err)
	}
}
