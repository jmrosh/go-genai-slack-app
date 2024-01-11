package slack

// Event is the format for slack events received
type Event struct {
	Token       string     `json:"token"`
	Challenge   string     `json:"challenge,omitempty"`
	TeamId      string     `json:"team_id,omitempty"`
	APIAppId    string     `json:"api_app_id,omitempty"`
	Type        string     `json:"type"`
	InnerEvent  InnerEvent `json:"event,omitempty"`
	AuthedTeams []string   `json:"authed_teams,omitempty"`
	EventId     string     `json:"event_id,omitempty"`
	EventTime   int        `json:"event_time,omitempty"`
}
