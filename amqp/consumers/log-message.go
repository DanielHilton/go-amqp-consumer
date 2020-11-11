package consumers

import (
	"encoding/json"
	"fmt"
	A "github.com/DanielHilton/go-amqp-consumer/amqp"
	"github.com/DanielHilton/go-amqp-consumer/structs"

	"github.com/streadway/amqp"
)

// NewLogMessageConsumer will create a channel and a consumer for the given queue name
func NewLogMessageConsumer(c *amqp.Connection, q string) {
	A.CreateConsumer(c, q, func(d amqp.Delivery) {
		var message structs.EnrichedMessage
		json.Unmarshal(d.Body, &message)

		fmt.Println("Loggy boi")
		d.Ack(false)
	})
}
