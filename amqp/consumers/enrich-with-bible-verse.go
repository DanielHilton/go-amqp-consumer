package consumers

import (
	"context"
	"encoding/json"
	"fmt"

	A "github.com/DanielHilton/go-amqp-consumer/amqp"
	"github.com/streadway/amqp"
	"go.mongodb.org/mongo-driver/mongo"
)

// NewEnrichWithBibleVerseConsumer will create a channel and a consumer for the given queue name and store the message to MongoDB
func NewEnrichWithBibleVerseConsumer(c *amqp.Connection, q string, client *mongo.Client) {
	A.CreateConsumer(c, q, func(d amqp.Delivery) {
		var message map[string]interface{}
		if msgErr := json.Unmarshal(d.Body, &message); msgErr != nil {
			fmt.Errorf("failed to process message %s for queue %s. Error: %w", string(d.Body), q, msgErr)
			d.Nack(false, false)
			return
		}

		fmt.Printf("Received message: %v\n", message)

		c := client.Database("poc").Collection("go")
		_, err := c.InsertOne(context.Background(), message)
		if err != nil {
			fmt.Errorf("failed to insert into DB")
			d.Nack(false, false)
		}

		d.Ack(false)
	})
}
