package consumers

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/DanielHilton/go-amqp-consumer/services"
	"github.com/DanielHilton/go-amqp-consumer/structs"
	"io/ioutil"
	"net/http"
	"time"

	A "github.com/DanielHilton/go-amqp-consumer/amqp"
	"github.com/streadway/amqp"
)

func getBiblePassage() structs.BiblePassage {
	r, err := http.Get("https://labs.bible.org/api/?type=json&passage=random")
	defer r.Body.Close()
	if err != nil {
		fmt.Errorf("failed to get bible passage")
		return structs.BiblePassage{}
	}
	fmt.Println("Got response from bible")

	bytes, _ := ioutil.ReadAll(r.Body)

	var passages []structs.BiblePassage
	jsonErr := json.Unmarshal(bytes, &passages)
	if jsonErr != nil {
		fmt.Errorf("failed to unmarshal %s", string(bytes))
	}

	return passages[0]
}

// NewEnrichWithBibleVerseConsumer will create a channel and a consumer for the given queue name and store the message to MongoDB
func NewEnrichWithBibleVerseConsumer(c *amqp.Connection, q string) {
	A.CreateConsumer(c, q, func(d amqp.Delivery) {
		var enrichedMessage structs.EnrichedMessage
		enrichedMessage.Timestamp = time.Now()

		if msgErr := json.Unmarshal(d.Body, &enrichedMessage.AMQPMessage); msgErr != nil {
			fmt.Errorf("failed to process enrichedMessage %s for queue %s. Error: %w", string(d.Body), q, msgErr)
			d.Nack(false, false)
			return
		}
		fmt.Println("Enriching enrichedMessage")

		enrichedMessage.BiblePassage = getBiblePassage()

		coll := services.MongoClient.Database("poc").Collection("go")
		_, err := coll.InsertOne(context.Background(), enrichedMessage)
		if err != nil {
			fmt.Errorf("failed to insert into DB")
			d.Nack(false, false)
		}

		bytes, _ := json.Marshal(enrichedMessage)
		A.PublishMessage(c, "test", "test.logmessage", bytes)
		d.Ack(false)
	})
}
