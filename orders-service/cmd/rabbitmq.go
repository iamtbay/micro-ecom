package main

import (
	"encoding/json"
	"fmt"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

func connectRabbitMQ() (*amqp.Channel, error) {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		return nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}
	return ch, nil
}

func consumeMessages(ch *amqp.Channel) {
	msgs, err := ch.Consume(
		"order_queue",
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("Failed to register a consumer %v", err)
	}

	go func() {
		for d := range msgs {
			log.Printf("Received message %s", d.Body)
			var orderCart Order
			err := json.Unmarshal([]byte(d.Body), &orderCart)
			if err != nil {
				log.Fatalf("Error unmarshalling JSON %v", err)
			}

			_, err = services.newOrder(orderCart)
			if err != nil {
				fmt.Println(err)
			}
			err = publishNotification(ch, MessageType{Message: "New order!"})
			if err != nil {
				fmt.Println(err)
			}
		}
	}()
}

func publishNotification(ch *amqp.Channel, msg MessageType) error {
	jsonData, err := json.Marshal(msg)
	if err != nil {
		return err
	}
	err = ch.Publish(
		"notifications_exchange",
		"order.created",
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        jsonData,
		},
	)
	if err != nil {
		return err
	}
	return nil

}
