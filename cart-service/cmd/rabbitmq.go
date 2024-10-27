package main

import (
	"encoding/json"
	"fmt"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

// CONNECT RABBITMQ
func connectRabbitMQ() (*amqp.Channel, error) {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		return nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	//define exchange
	err = ch.ExchangeDeclare(
		"cart_exchange",
		"direct",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return nil, err
	}

	//define queue
	_, err = ch.QueueDeclare(
		"order_queue",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return nil, err
	}

	//bind queue to exchange

	err = ch.QueueBind(
		"order_queue",
		"cart_confirmed",
		"cart_exchange",
		false,
		nil,
	)
	if err != nil {
		return nil, err
	}
	return ch, nil
}

func publishMessage(ch *amqp.Channel, order CartOrder) error {
	jsonData, err := json.Marshal(order)
	if err != nil {
		return err
	}

	err = ch.Publish(
		"cart_exchange",
		"cart_confirmed",
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        jsonData,
		})
	if err != nil {
		return err
	}
	return nil
}



func publishNotification(ch *amqp.Channel, msg MessageType) error {
	jsonData, err := json.Marshal(msg)
	if err != nil {
		return err
	}
	err = ch.Publish(
		"notifications_exchange",
		"cart.updated",
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

func consumeMessages() {
	msgs, err := ch.Consume(
		"price_updates_queue",
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
		fmt.Println("listening for changes....")
		for d := range msgs {
			var product UpdateProductType
			err := json.Unmarshal(d.Body, &product)
			if err != nil {
				log.Fatal("something went wrong while listening changes", err)
			}
			err = services.updateProduct(product)
			if err != nil {
				log.Fatal("something went wrong while listening changes", err)
			}
			//notification
			err = publishNotification(ch, MessageType{Message: "Price has changed"})
			if err != nil {
				log.Fatal("something went wrong while sending notification", err)
			}
			fmt.Println("consumer did his work!")
		}
	}()
}