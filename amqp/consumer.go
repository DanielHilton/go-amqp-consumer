package amqp

import (
	"fmt"
	H "github.com/DanielHilton/go-amqp-consumer/helpers"
	"github.com/streadway/amqp"
	"time"
)

func processDelivery(messages <-chan amqp.Delivery, handler func(delivery amqp.Delivery, t chan time.Time)) {
	for m := range messages {
		now := time.Now()
		timerChan := make(chan time.Time, 1)
		go handler(m, timerChan)

		done := <-timerChan
		elapsed := done.Sub(now)

		fmt.Printf("Message processed: %s\n", elapsed.String())
	}
}

func CreateConsumer(c *amqp.Connection, q string, handler func(amqp.Delivery, chan time.Time)) {
	ch, err := c.Channel()
	H.ExitOnFail(err, "Failed to create a channel for consumer")

	msgChan, err := ch.Consume(q, "", false, false, false, false, nil)
	H.ExitOnFail(err, "Failed to register consumer for queue")

	go processDelivery(msgChan, handler)
}
