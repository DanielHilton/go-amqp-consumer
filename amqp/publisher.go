package amqp

import (
	"fmt"
	"time"

	AMQP "github.com/streadway/amqp"
)

func PublishMessage(c *AMQP.Connection, x string, rk string, b []byte) {
	ch, chErr := c.Channel()
	if chErr != nil {
		fmt.Error("failed to create channel for publisher")
		return
	}
	defer ch.Close()

	msg := AMQP.Publishing{
		DeliveryMode: AMQP.Persistent,
		Timestamp:    time.Now(),
		ContentType:  "application/json",
		Body:         b,
	}

	pubErr := ch.Publish(x, rk, true, true, msg)
	if pubErr != nil {
		fmt.Errorf("failed to publish message onto exchange %s, routing key %s", x, rk)
	}
}
