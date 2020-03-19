package consumers

import (
	"encoding/json"
	"log"
	"time"

	H "github.com/DanielHilton/go-amqp-consumer/helpers"
	"github.com/streadway/amqp"
)

// Create will create a channel and a consumer for the given queue name
func Create(c *amqp.Connection, q string) {
	ch, err := c.Channel()
	H.FailOnError(err, "Failed to create a channel for consumer")

	msgs, err := ch.Consume(q, "", true, false, false, false, nil)
	H.FailOnError(err, "Failed to register consumer for queue")

	forever := make(chan bool)

	go func() {
		for m := range msgs {
			b, err := json.Marshal(struct {
				x time.Time
				y amqp.Delivery
			}{x: time.Now(), y: m})

			if err != nil {
				log.Println((b))
			} else {
				log.Printf("%s", err)
			}
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
