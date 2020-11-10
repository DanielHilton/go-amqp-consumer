package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"

	C "github.com/DanielHilton/go-amqp-consumer/amqp/consumers"
	H "github.com/DanielHilton/go-amqp-consumer/helpers"
	"github.com/streadway/amqp"
)

func main() {
	mongoClient, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = mongoClient.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer mongoClient.Disconnect(ctx)

	exchange := "test"
	amqpUri := "amqp://guest:guest@localhost:5672"
	conn, err := amqp.Dial(amqpUri)
	H.FailOnError(err, "Failed to open a connection")
	fmt.Printf("Connected to %s successfully\n", amqpUri)
	defer conn.Close()

	ch, err := conn.Channel()
	H.FailOnError(err, "Failed to open a channel")
	defer ch.Close()

	err = ch.ExchangeDeclare(
		exchange, "topic",
		false, false, false, false, nil)
	H.FailOnError(err, "Failed to declare exchange")
	fmt.Printf("Exchange %s declared\n", exchange)

	_, err = ch.QueueDeclare(
		"enrichWithBibleVerse", // name
		false,                  // durable
		false,                  // delete when unused
		false,                  // exclusive
		false,                  // no-wait
		nil,                    // arguments
	)
	H.FailOnError(err, "Failed to create queue")
	fmt.Printf("Queue %s declared\n", "enrichWithBibleVerse")

	_, err = ch.QueueDeclare(
		"logMessage", // name
		false,        // durable
		false,        // delete when unused
		false,        // exclusive
		false,        // no-wait
		nil,          // arguments
	)
	H.FailOnError(err, "Failed to create queue")
	fmt.Printf("Queue %s declared\n", "logMessage")

	err = ch.QueueBind("enrichWithBibleVerse", "test.enrichwithbibleverse", "test", false, nil)
	err = ch.QueueBind("logMessage", "test.logmessage", "test", false, nil)
	H.FailOnError(err, "Failed to bind queue to exchange")

	C.CreateStoreMongoConsumer(conn, "enrichWithBibleVerse", mongoClient)
	C.CreatePrintConsumer(conn, "logMessage")

	forever := make(chan bool)
	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
