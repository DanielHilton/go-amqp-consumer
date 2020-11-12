package main

import (
	"context"
	"fmt"
	"github.com/DanielHilton/go-amqp-consumer/services"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	A "github.com/DanielHilton/go-amqp-consumer/amqp"
	C "github.com/DanielHilton/go-amqp-consumer/amqp/consumers"
	H "github.com/DanielHilton/go-amqp-consumer/helpers"
	"github.com/streadway/amqp"
)

func main() {
	mc, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}

	services.MongoClient = mc

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	err = services.MongoClient.Connect(ctx)
	H.ExitOnFail(err, "failed to connect to MongoDB")
	defer cancel()
	defer services.MongoClient.Disconnect(ctx)

	amqpURI := "amqp://guest:guest@localhost:5672"
	conn, err := amqp.Dial(amqpURI)
	H.ExitOnFail(err, "Failed to open a connection")
	defer conn.Close()

	fmt.Printf("Connected to %s successfully\n", amqpURI)

	A.CreateExchange(conn, "test", "topic")
	A.CreateQueue(conn, "enrichWithBibleVerse")
	A.CreateQueue(conn, "logMessage")

	{
		ch, _ := conn.Channel()
		defer ch.Close()

		err = ch.QueueBind("enrichWithBibleVerse", "test.enrichwithbibleverse", "test", false, nil)
		H.ExitOnFail(err, "Failed to bind queue to exchange")
		err = ch.QueueBind("logMessage", "test.logmessage", "test", false, nil)
		H.ExitOnFail(err, "Failed to bind queue to exchange")
	}

	C.NewEnrichWithBibleVerseConsumer(conn, "enrichWithBibleVerse")
	C.NewLogMessageConsumer(conn, "logMessage")

	H.StartHttpServer(9000)

	forever := make(chan bool)
	log.Printf("The Golang POC has started")
	<-forever
}
