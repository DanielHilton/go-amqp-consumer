package amqp

import (
	H "github.com/DanielHilton/go-amqp-consumer/helpers"
	"github.com/streadway/amqp"
)

func processDelivery(messages <-chan amqp.Delivery, handler func(delivery amqp.Delivery)) {
	for m := range messages {
		go handler(m)
	}
}

func CreateConsumer(c *amqp.Connection, q string, handler func(amqp.Delivery)) {
	ch, err := c.Channel()
	H.FailOnError(err, "Failed to create a channel for consumer")

	msgChan, err := ch.Consume(q, "", false, false, false, false, nil)
	H.FailOnError(err, "Failed to register consumer for queue")

	go processDelivery(msgChan, handler)
}
