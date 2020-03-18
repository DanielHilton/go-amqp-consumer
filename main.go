package main

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

func main() {
	qname := "go-stuff"
	xname := "test"

	uri := "amqp://guest:guest@localhost:5672/test"
	conn, err := amqp.Dial(uri)
	failOnError(err, "Failed to open a connection")
	fmt.Printf("Connected to %s successfully\n", uri)
	defer conn.Close()

	channel, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer channel.Close()

	err = channel.ExchangeDeclare(
		xname, "fanout",
		true, false, false, false, nil)
	failOnError(err, "Failed to declare exchange")
	fmt.Printf("Exchange %s declared\n", xname)

	_, err = channel.QueueDeclare(
		qname, // name
		false, // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	failOnError(err, "Failed to create queue")
	fmt.Printf("Queue %s declared\n", qname)

	err = channel.QueueBind("go-stuff", "test.go-stuff", "test", false, nil)
	failOnError(err, "Failed to bind queue to exchange")
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s, %s", msg, err)
	}
}
