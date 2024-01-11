package firestore

type ConversationDto struct {
	Messages MessageContainer `json:"messages"`
}

type MessageContainer struct {
	ArrayValue MessageArrayContainer `json:"arrayValue"`
}

type MessageArrayContainer struct {
	Values []MessageValue `json:"values"`
}

type MessageValue struct {
	MapValue MessageMapValue `json:"mapValue"`
}

type MessageMapValue struct {
	Fields MessageFields `json:"fields"`
}

type MessageFields struct {
	Role RoleValue `json:"role"`
	Text TextValue `json:"text"`
}

type RoleValue struct {
	StringValue string `json:"stringValue"`
}

type TextValue struct {
	StringValue string `json:"stringValue"`
}
