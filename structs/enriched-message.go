package structs

import "time"

type EnrichedMessage struct {
	AmqpMessage  interface{}
	Timestamp    time.Time
	BiblePassage BiblePassage
}
