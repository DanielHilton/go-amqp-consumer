package consumers

import (
	"context"
	"encoding/json"
	"fmt"
	A "github.com/DanielHilton/go-amqp-consumer/amqp"
	"github.com/streadway/amqp"
	"go.mongodb.org/mongo-driver/mongo"
)

// CreateStoreMongoConsumer will create a channel and a consumer for the given queue name and store the message to MongoDB
func CreateStoreMongoConsumer(c *amqp.Connection, q string, client *mongo.Client) {
	A.CreateConsumer(c, q, func(messages <-chan amqp.Delivery) {
		for m := range messages {
			var message map[string]interface{}
			if msgErr := json.Unmarshal(m.Body, &message); msgErr != nil {
				fmt.Errorf("failed to process message %s for queue %s. Error: %v", string(m.Body), q, msgErr)
				m.Nack(false, false)
				return
			}

			fmt.Printf("Received message: %v\n", message)

			c := client.Database("poc").Collection("go")
			_, err := c.InsertOne(context.Background(), message)
			if err != nil {
				fmt.Errorf("failed to insert into DB")
				m.Nack(false, false)
			}

			m.Ack(false)
		}
	})
}
