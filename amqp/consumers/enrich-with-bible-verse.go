package consumers

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/DanielHilton/go-amqp-consumer/structs"
	"io/ioutil"
	"net/http"
	"time"

	A "github.com/DanielHilton/go-amqp-consumer/amqp"
	"github.com/streadway/amqp"
	"go.mongodb.org/mongo-driver/mongo"
)

func getBiblePassage() structs.BiblePassage {
	r, err := http.Get("https://labs.bible.org/api/?type=json&passage=random")
	defer r.Body.Close()
	if err != nil {
		fmt.Errorf("failed to get bible passage")
		return structs.BiblePassage{}
	}

	bytes, _ := ioutil.ReadAll(r.Body)

	var passages []structs.BiblePassage
	jsonErr := json.Unmarshal(bytes, &passages)
	if jsonErr != nil {
		fmt.Errorf("failed to unmarshal %s", string(bytes))
	}

	return passages[0]
}

// NewEnrichWithBibleVerseConsumer will create a channel and a consumer for the given queue name and store the message to MongoDB
func NewEnrichWithBibleVerseConsumer(c *amqp.Connection, q string, mc *mongo.Client) {
	A.CreateConsumer(c, q, func(d amqp.Delivery) {
		var message structs.EnrichedMessage
		message.Timestamp = time.Now()

		if msgErr := json.Unmarshal(d.Body, &message.AmqpMessage); msgErr != nil {
			fmt.Errorf("failed to process message %s for queue %s. Error: %w", string(d.Body), q, msgErr)
			d.Nack(false, false)
			return
		}
		fmt.Println("Enriching message")

		message.BiblePassage = getBiblePassage()

		coll := mc.Database("poc").Collection("go")
		_, err := coll.InsertOne(context.Background(), message)
		if err != nil {
			fmt.Errorf("failed to insert into DB")
			d.Nack(false, false)
		}

		A.PublishMessage(c, "test", "test.logmessage", d.Body)
		d.Ack(false)
	})
}
