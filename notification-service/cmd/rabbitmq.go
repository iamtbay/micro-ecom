package main

import (
	"log"

	"github.com/gorilla/websocket"
	amqp "github.com/rabbitmq/amqp091-go"
)

func connectRabbitMQ() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ %s", err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("failed to open a channel %s", err)
	}
	defer ch.Close()

	err = ch.ExchangeDeclare(
		"notifications_exchange",
		"topic",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("Failed to declare exchange %s", err)
	}

	q, err := ch.QueueDeclare(
		"notification_queue",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("Failed to declare a queue %s", err)
	}
	routingKeys := []string{"order.*", "cart.*"}
	for _, routingKey := range routingKeys {
		err := ch.QueueBind(
			q.Name,
			routingKey,
			"notifications_exchange",
			false,
			nil,
		)
		if err != nil {
			log.Fatalf("failed to bind queue %s", err)
		}
	}

	//consume
	consumeMessages(ch)
}

func consumeMessages(ch *amqp.Channel) {
	msgs, err := ch.Consume(
		"notification_queue",
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("Failed to register a consumer %s", err)
	}
	for msg := range msgs {
		log.Printf("Received notification %s", msg.Body)
		sendToClients(string(msg.Body))
	}
}

func sendToClients(message string) {
	mu.Lock()
	defer mu.Unlock()
	for client := range clients {
		err := client.WriteMessage(websocket.TextMessage, []byte(message))
		if err != nil {
			log.Println("Error sendind message to client:", err)
			client.Close()
			delete(clients, client)
		}
	}
}
