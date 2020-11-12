package amqp

import (
	"fmt"
	H "github.com/DanielHilton/go-amqp-consumer/helpers"
	"github.com/streadway/amqp"
)

// CreateQueue is a convenience function that creates a simple QueueDeclare with some default settings
// A non-durable, auto-delete disabled, non-exclusive, waiting Queue with no arguments
func CreateQueue(c *amqp.Connection, name string) amqp.Queue {
	ch, _ := c.Channel()
	defer ch.Close()

	q, err := ch.QueueDeclare(name, false, false, false, false, nil)
	H.ExitOnFail(err, "Queue declaration failed.")

	fmt.Printf("%s Queue declared\n", name)
	return q
}

func CreateExchange(c *amqp.Connection, name string, kind string) {
	ch, _ := c.Channel()
	defer ch.Close()

	err := ch.ExchangeDeclare(name, kind, false, false, false, false, nil)
	H.ExitOnFail(err, "Failed to create exchange")
}
