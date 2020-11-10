package amqp

import (
	H "github.com/DanielHilton/go-amqp-consumer/helpers"
	"github.com/streadway/amqp"
)

func CreateConsumer(c *amqp.Connection, q string, f func(<-chan amqp.Delivery)) {
	ch, err := c.Channel()
	H.FailOnError(err, "Failed to create a channel for consumer")

	msgChan, err := ch.Consume(q, "", false, false, false, false, nil)
	H.FailOnError(err, "Failed to register consumer for queue")

	go f(msgChan)
}
