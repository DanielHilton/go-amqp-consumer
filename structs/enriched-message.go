package structs

import (
	"encoding/json"
	"time"
)

type EnrichedMessage struct {
	AMQPMessage  interface{}   `json:"amqpMessage"`
	Timestamp    time.Time     `json:"timestamp"`
	BiblePassage *BiblePassage `json:"biblePassage"`
}

// A "String() string" method for EnrichedMessage.
func (a EnrichedMessage) String() string {
	bytes, _ := json.Marshal(a)

	return string(bytes)
}
