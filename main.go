package main

import (
	"fmt"
	A "github.com/DanielHilton/go-amqp-consumer/amqp"
	C "github.com/DanielHilton/go-amqp-consumer/amqp/consumers"
	H "github.com/DanielHilton/go-amqp-consumer/helpers"
	"github.com/DanielHilton/go-amqp-consumer/services/db"
	"github.com/DanielHilton/go-amqp-consumer/services/server"
	"github.com/streadway/amqp"
	"log"
)

func main() {
	ctx, cancel := db.Init()
	defer cancel()
	defer db.MongoClient.Disconnect(ctx)

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

	server.StartHttpServer(9000)

	forever := make(chan bool)
	log.Printf("The Golang POC has started")
	<-forever
}
