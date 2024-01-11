package firestore

// Message is the format of a message sent to/from slack
type Message struct {
	Role string `json:"role"`
	Text string `json:"text"`
	//TS   string `json:"ts"`
}
