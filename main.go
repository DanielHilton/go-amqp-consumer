package main

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

func main() {
	uri := "amqp://guest:guest@localhost:5672/test"
	conn, err := amqp.Dial(uri)
	failOnError(err, "Failed to open a connection")
	fmt.Printf("Connected to %s successfully", uri)
	defer conn.Close()

	channel, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer channel.Close()
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s, %s", msg, err)
	}
}
