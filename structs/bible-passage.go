package structs

// BiblePassage is a struct which represents a passage response from the Bible API
type BiblePassage struct {
	Book    string `json:"bookname"`
	Chapter string
	Verse   string
	Text    string
}
