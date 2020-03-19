package consumers

import (
	"fmt"
	"log"
	"time"

	H "github.com/DanielHilton/go-amqp-consumer/helpers"
	"github.com/streadway/amqp"
)

func processMessages(msgs <-chan amqp.Delivery) {
	for d := range msgs {
		l := struct {
			t int64
			b string
		}{t: time.Now().Unix(), b: string(d.Body)}

		fmt.Printf("Received message: %v\n", l)
	}
}

// Create will create a channel and a consumer for the given queue name
func Create(c *amqp.Connection, q string) {
	ch, err := c.Channel()
	H.FailOnError(err, "Failed to create a channel for consumer")

	msgs, err := ch.Consume(q, "", true, false, false, false, nil)
	H.FailOnError(err, "Failed to register consumer for queue")

	forever := make(chan bool)

	go processMessages(msgs)

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
