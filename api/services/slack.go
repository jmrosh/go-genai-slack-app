package services

import (
	"fmt"
	"github.com/slack-go/slack"
)

type SlackService interface {
	SendMessage(text string, channel string) error
}

type SlackClient interface {
	PostMessage(channelID string, options ...slack.MsgOption) (string, string, error)
}

// SlackService is a struct that holds the Slack service functionalities.
type slackService struct {
	client SlackClient
}

// NewSlackService is a constructor function for SlackService.
func NewSlackService(client SlackClient) SlackService {
	return &slackService{
		client: client,
	}
}

// SendMessage sends a message on behalf of a Slack app.
func (s *slackService) SendMessage(text string, channel string) error {
	if text == "" {
		return fmt.Errorf("SendMessage(): text cannot be empty\n")
	}
	if channel == "" {
		return fmt.Errorf("SendMessage(): channel cannot be empty\n")
	}

	_, _, err := s.client.PostMessage(
		channel,
		slack.MsgOptionText(text, false),
		//slack.MsgOptionAttachments(attachment),
	)

	if err != nil {
		return fmt.Errorf("\n\nPostMessage(). Error: %v\n", err)
	}

	return nil
}
