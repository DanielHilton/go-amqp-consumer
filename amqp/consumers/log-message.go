package consumers

import (
	"fmt"
	"time"

	A "github.com/DanielHilton/go-amqp-consumer/amqp"

	"github.com/streadway/amqp"
)

// NewLogMessageConsumer will create a channel and a consumer for the given queue name
func NewLogMessageConsumer(c *amqp.Connection, q string) {
	A.CreateConsumer(c, q, func(d amqp.Delivery) {
		l := struct {
			t int64
			b string
		}{t: time.Now().Unix(), b: string(d.Body)}

		fmt.Printf("Received message: %v\n", l)
		d.Ack(false)
	})
}
