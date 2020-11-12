package consumers

import (
	"encoding/json"
	"fmt"
	"github.com/DanielHilton/go-amqp-consumer/db"
	"github.com/DanielHilton/go-amqp-consumer/structs"
	"io/ioutil"
	"net/http"
	"time"

	A "github.com/DanielHilton/go-amqp-consumer/amqp"
	"github.com/streadway/amqp"
)

func getBiblePassage() (*structs.BiblePassage, error) {
	r, err := http.Get("https://labs.bible.org/api/?type=json&passage=random")
	defer r.Body.Close()
	if err != nil {
		fmt.Errorf("failed to get bible passage")
		return nil, err
	}
	fmt.Println("Got response from bible")

	bytes, _ := ioutil.ReadAll(r.Body)
	fmt.Print(string(bytes))

	var passages []structs.BiblePassage
	err = json.Unmarshal(bytes, &passages)
	if err != nil {
		fmt.Errorf("failed to unmarshal %s", string(bytes))
		return nil, err
	}

	return &passages[0], nil
}

// NewEnrichWithBibleVerseConsumer will create a channel and a consumer for the given queue name and store the message to MongoDB
func NewEnrichWithBibleVerseConsumer(c *amqp.Connection, q string) {
	A.CreateConsumer(c, q, func(d amqp.Delivery, t chan time.Time) {
		var enrichedMessage structs.EnrichedMessage
		enrichedMessage.Timestamp = time.Now()

		if msgErr := json.Unmarshal(d.Body, &enrichedMessage.AMQPMessage); msgErr != nil {
			fmt.Errorf("failed to process enrichedMessage %s for queue %s. Error: %w", string(d.Body), q, msgErr)
			d.Nack(false, false)
			return
		}

		bp, err := getBiblePassage()
		if err != nil {
			fmt.Errorf("failed to get bible passage %w", err)
			d.Nack(false, false)
			t <- time.Now()
			return
		}

		enrichedMessage.BiblePassage = bp
		err = db.InsertEnrichedMessage(enrichedMessage)
		if err != nil {
			fmt.Errorf("failed to insert into DB %w", err)
			d.Nack(false, false)
			t <- time.Now()
			return
		}

		bytes, _ := json.Marshal(enrichedMessage)
		A.PublishMessage(c, "test", "test.logmessage", bytes)
		d.Ack(false)

		t <- time.Now()
	})
}
