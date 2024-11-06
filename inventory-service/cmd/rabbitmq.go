package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/rabbitmq/amqp091-go"
)

func connectRabbitMQ() {
	conn, err := amqp091.Dial("amqp://guest:guest@localhost:5672/")
	failOnErr("error while connecting", err)

	ch, err := conn.Channel()
	failOnErr("error while opening channel", err)

	err = ch.ExchangeDeclare(
		"inventory_exchange",
		"topic",
		true,
		false,
		false,
		false,
		nil,
	)
	failOnErr("failed while declaring an exchange", err)

	queues := []string{"inventory_reserve_queue", "inventory_cancel_queue", "inventory_sold_queue", "inventory_new_product_queue", "inventory_check_queue"}
	for _, queue := range queues {
		_, err := ch.QueueDeclare(
			queue,
			true,
			false,
			false,
			false,
			nil,
		)
		failOnErr("error while declare a queue", err)
	}
	queueBindings := map[string]string{
		"inventory_reserve_queue":     "inventory.reserve",
		"inventory_cancel_queue":      "inventory.cancel",
		"inventory_sold_queue":        "inventory.sold",
		"inventory_new_product_queue": "inventory.new",
		"inventory_check_queue":       "inventory.check",
	}
	for queue, key := range queueBindings {

		err = ch.QueueBind(
			queue,
			key,
			"inventory_exchange",
			false,
			nil,
		)

		failOnErr("failed while binding queue", err)
	}
	//
	go consumeReserveMessages(ch)
	go consumeCancelMessages(ch)
	go consumeSoldMessages(ch)
	go consumeNewProductMessages(ch)
	go consumeCheckMessage(ch)

}

// !
// RESERVE
func consumeReserveMessages(ch *amqp091.Channel) {
	msgs, err := ch.Consume(
		"inventory_reserve_queue",
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	failOnErr("error while consuming messages", err)

	for msg := range msgs {
		var productData ProductData
		err := json.Unmarshal(msg.Body, &productData)
		failOnErr("err marshalling data", err)

		err = services.updateStockViaReserved(productData)
		failOnErr("err stocking data", err)

		log.Println("Inventory updated via reserved!")
	}
}

// !
// CANCEL
func consumeCancelMessages(ch *amqp091.Channel) {
	msgs, err := ch.Consume(
		"inventory_cancel_queue",
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	failOnErr("consume cancel messages", err)

	for msg := range msgs {
		var productData ProductData
		err := json.Unmarshal(msg.Body, &productData)
		failOnErr("err un marshaling", err)

		err = services.cancelReservation(productData)
		failOnErr("err cancelling reservation", err)

		log.Println("Inventory updated via canceled!")
	}
}

// !
// SOLD
func consumeSoldMessages(ch *amqp091.Channel) {
	msgs, err := ch.Consume(
		"inventory_sold_queue",
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	failOnErr("err consume sold messages", err)

	for msg := range msgs {
		var productData ProductData
		err := json.Unmarshal(msg.Body, &productData)
		failOnErr("something went wrong marshaling", err)

		err = services.updateStockViaSold(productData)
		failOnErr("something went wrong", err)

		log.Println("Inventory updated via sold!")
	}
}

// !
// NEW
func consumeNewProductMessages(ch *amqp091.Channel) {
	msgs, err := ch.Consume(
		"inventory_new_product_queue",
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	failOnErr("err while consuming new prod msgs", err)

	for msg := range msgs {
		var productData ProductData
		err = json.Unmarshal(msg.Body, &productData)
		failOnErr("err unmarshaling", err)

		err = services.newProductStock(Product{
			ProductID:      productData.ProductID,
			AvailableStock: int64(productData.Quantity),
		})
		failOnErr("error while adding new product", err)

		fmt.Println("new product stock added")
	}
}

func failOnErr(msg string, err error) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func consumeCheckMessage(ch *amqp091.Channel) {
	msgs, err := ch.Consume(
		"inventory_check_queue",
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("Failed to consume check messages: %s", err)
	}

	for msg := range msgs {
		stockAvailable, err := services.checkStock(string(msg.Body))
		if err != nil {
			log.Fatalf("Failed to consume checkstock messages1: %s", err)
		}

		err = ch.Publish(
			"",
			msg.ReplyTo,
			false,
			false,
			amqp091.Publishing{
				ContentType:   "text/plain",
				CorrelationId: msg.CorrelationId,
				Body:          stockAvailable,
			},
		)
		if err != nil {
			log.Fatalf("Failed to consume checkstock messages: %s", err)
		}

	}
}
