package firestore

// EventDto is the payload of a Firestore event.
// Refer to the docs for additional information regarding Firestore events.
type EventDto struct {
	//OldValue Value `json:"oldValue"`
	Value Value `json:"value"`
}

// Value holds Firestore document fields
type Value struct {
	Name   string          `json:"name"`
	Fields ConversationDto `json:"fields"`
}
