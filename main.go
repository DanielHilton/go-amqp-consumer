package main

import (
	"fmt"
	"log"

	C "github.com/DanielHilton/go-amqp-consumer/consumers"
	H "github.com/DanielHilton/go-amqp-consumer/helpers"
	"github.com/streadway/amqp"
)

func main() {
	exchange := "test"

	uri := "amqp://guest:fishcake@localhost:5672/test"
	conn, err := amqp.Dial(uri)
	H.FailOnError(err, "Failed to open a connection")
	fmt.Printf("Connected to %s successfully\n", uri)
	defer conn.Close()

	ch, err := conn.Channel()
	H.FailOnError(err, "Failed to open a channel")
	defer ch.Close()

	err = ch.ExchangeDeclare(
		exchange, "fanout",
		true, false, false, false, nil)
	H.FailOnError(err, "Failed to declare exchange")
	fmt.Printf("Exchange %s declared\n", exchange)

	_, err = ch.QueueDeclare(
		"go-stuff", // name
		false,      // durable
		false,      // delete when unused
		false,      // exclusive
		false,      // no-wait
		nil,        // arguments
	)
	H.FailOnError(err, "Failed to create queue")
	fmt.Printf("Queue %s declared\n", "go-stuff")

	_, err = ch.QueueDeclare(
		"mongo-stuff", // name
		false,         // durable
		false,         // delete when unused
		false,         // exclusive
		false,         // no-wait
		nil,           // arguments
	)
	H.FailOnError(err, "Failed to create queue")
	fmt.Printf("Queue %s declared\n", "mongo-stuff")

	err = ch.QueueBind("go-stuff", "test.go-stuff", "test", false, nil)
	err = ch.QueueBind("mongo-stuff", "test.mongo-stuff", "test", false, nil)
	H.FailOnError(err, "Failed to bind queue to exchange")

	C.CreatePrintConsumer(conn, "go-stuff")
	C.CreateStoreMongoConsumer(conn, "mongo-stuff")

	forever := make(chan bool)
	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
