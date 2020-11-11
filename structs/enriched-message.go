package structs

import (
	"encoding/json"
	"time"
)

type EnrichedMessage struct {
	AMQPMessage  interface{}
	Timestamp    time.Time
	BiblePassage BiblePassage
}

// A "String() string" method for EnrichedMessage.
func (a EnrichedMessage) String() string {
	bytes, _ := json.Marshal(a)

	return string(bytes)
}
