package structs

import (
	"encoding/json"
)

// BiblePassage is a struct which represents a passage response from the Bible API
type BiblePassage struct {
	Book    string `json:"bookname"`
	Chapter string `json:"chapter"`
	Verse   string `json:"verse"`
	Text    string `json:"text"`
}

func (b BiblePassage) String() string {
	bytes, _ := json.Marshal(b)

	return string(bytes)
}
