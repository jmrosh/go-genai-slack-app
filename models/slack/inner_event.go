package slack

const (
	MessageEvent         = "message"
	MessageAppHomeEvent  = "message.app_home"
	MessageChannelsEvent = "message.channels"
	MessageGroupsEvent   = "message.groups"
	MessageIMEvent       = "message.im"
	MessageMPIMEvent     = "message.mpim"
)

// InnerEvent is the format for content of slack events received
type InnerEvent struct {
	Type        string `json:"type"`
	BotId       string `json:"bot_id"`
	Channel     string `json:"channel"`
	User        string `json:"user"`
	Text        string `json:"text"`
	TS          string `json:"ts"`
	EventTS     string `json:"event_ts"`
	ChannelType string `json:"channel_type"`
}
