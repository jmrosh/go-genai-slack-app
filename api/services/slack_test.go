package services

import (
	"github.com/jmrosh/go-genai-slack-app/api/services/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSlackService(t *testing.T) {
	t.Run("Valid input", func(t *testing.T) {
		slackSvc := &slackService{client: &mocks.MockSlackClient{}}
		err := slackSvc.SendMessage("Hello, World!", "#channel")
		assert.Nil(t, err, "Error should be nil")
	})

	t.Run("Empty text", func(t *testing.T) {
		slackSvc := &slackService{client: &mocks.MockSlackClient{}}
		err := slackSvc.SendMessage("", "#channel")
		assert.NotNil(t, err, "Error should not be nil")
	})

	t.Run("Empty channel", func(t *testing.T) {
		slackSvc := &slackService{client: &mocks.MockSlackClient{}}
		err := slackSvc.SendMessage("Hello, World!", "")
		assert.NotNil(t, err, "Error should not be nil")
	})

	t.Run("Both empty", func(t *testing.T) {
		slackSvc := &slackService{client: &mocks.MockSlackClient{}}
		err := slackSvc.SendMessage("", "")
		assert.NotNil(t, err, "Error should not be nil")
	})
}
