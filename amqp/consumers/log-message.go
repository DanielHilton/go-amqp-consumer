package consumers

import (
	"fmt"
	"time"

	H "github.com/DanielHilton/go-amqp-consumer/helpers"
	"github.com/streadway/amqp"
)

// CreatePrintConsumer will create a channel and a consumer for the given queue name
func CreatePrintConsumer(c *amqp.Connection, q string) {
	ch, err := c.Channel()
	H.FailOnError(err, "Failed to create a channel for consumer")

	messages, err := ch.Consume(q, "", true, false, false, false, nil)
	H.FailOnError(err, "Failed to register consumer for queue")

	go func(messages <-chan amqp.Delivery) {
		for d := range messages {
			l := struct {
				t int64
				b string
			}{t: time.Now().Unix(), b: string(d.Body)}

			fmt.Printf("Received message: %v\n", l)
		}
	}(messages)
}
