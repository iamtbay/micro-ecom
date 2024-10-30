package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/rabbitmq/amqp091-go"
)

var ch *amqp091.Channel

func connectRabbitMQ() {
	var err error
	conn, err := amqp091.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatalf("Failed to declare queue: %v", err)
	}

	ch, err = conn.Channel()
	if err != nil {
		log.Fatalf("Failed to declare queue: %v", err)
	}

	err = ch.ExchangeDeclare(
		"price_updates_exchange",
		"direct",
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		log.Fatalf("Failed to declare queue: %v", err)
	}

	//declare
	priceQueue, err := ch.QueueDeclare(
		"price_updates_queue",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("Failed to declare queue: %v", err)
	}
	//bind
	err = ch.QueueBind(
		priceQueue.Name,
		"",
		"price_updates_exchange",
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("Faield to bind price_updates_queue %v", err)
	}
}

func publishPrice(product GetProduct) error {
	jsonData, err := json.Marshal(product)
	if err != nil {
		return err
	}

	err = ch.Publish(
		"price_updates_exchange",
		"",
		false,
		false,
		amqp091.Publishing{
			ContentType: "application/json",
			Body:        jsonData,
		})
	if err != nil {
		return err
	}
	fmt.Println("published!")
	return nil
}

func publishNewProduct(productInventory ProductInventoryType) error {

	jsonData, err := json.Marshal(&productInventory)
	if err != nil {
		return err
	}
	err = ch.Publish(
		"inventory_exchange",
		"inventory.new",
		false,
		false,
		amqp091.Publishing{
			ContentType: "application/json",
			Body:        jsonData,
		},
	)
	if err != nil {
		return err
	}
	return nil

}
