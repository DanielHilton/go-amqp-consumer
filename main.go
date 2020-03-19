package main

import (
	"fmt"

	printConsumer "github.com/DanielHilton/go-amqp-consumer/consumers"
	H "github.com/DanielHilton/go-amqp-consumer/helpers"
	"github.com/streadway/amqp"
)

func main() {
	q := "go-stuff"
	x := "test"

	uri := "amqp://guest:guest@localhost:5672/test"
	conn, err := amqp.Dial(uri)
	H.FailOnError(err, "Failed to open a connection")
	fmt.Printf("Connected to %s successfully\n", uri)
	defer conn.Close()

	ch, err := conn.Channel()
	H.FailOnError(err, "Failed to open a channel")
	defer ch.Close()

	err = ch.ExchangeDeclare(
		x, "fanout",
		true, false, false, false, nil)
	H.FailOnError(err, "Failed to declare exchange")
	fmt.Printf("Exchange %s declared\n", x)

	_, err = ch.QueueDeclare(
		q,     // name
		false, // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	H.FailOnError(err, "Failed to create queue")
	fmt.Printf("Queue %s declared\n", q)

	err = ch.QueueBind("go-stuff", "test.go-stuff", "test", false, nil)
	H.FailOnError(err, "Failed to bind queue to exchange")

	printConsumer.Create(conn, "go-stuff")
}
