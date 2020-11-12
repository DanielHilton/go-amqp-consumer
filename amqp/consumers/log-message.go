package consumers

import (
	"fmt"
	A "github.com/DanielHilton/go-amqp-consumer/amqp"
	"github.com/streadway/amqp"
	"time"
)

// NewLogMessageConsumer will create a channel and a consumer for the given queue name
func NewLogMessageConsumer(c *amqp.Connection, q string) {
	A.CreateConsumer(c, q, func(d amqp.Delivery, t chan time.Time) {

		fmt.Println(string(d.Body))
		d.Ack(false)

		t <- time.Now()
	})
}
