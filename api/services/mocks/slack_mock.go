package mocks

import (
	"fmt"
	"github.com/slack-go/slack"
	"github.com/stretchr/testify/mock"
)

type MockSlackClient struct {
	mock.Mock
}

type MockSlackService struct {
	mock.Mock
}

func NewMockSlackClient() *MockSlackClient {
	return &MockSlackClient{}
}

func NewMockSlackService() *MockSlackService {
	return &MockSlackService{}
}

func (msc *MockSlackClient) PostMessage(channelID string, options ...slack.MsgOption) (string, string, error) {
	if channelID == "" {
		return "", "", fmt.Errorf("input cannot be empty")
	}
	return "", "", nil
}

func (m *MockSlackService) SendMessage(text string, channel string) error {
	args := m.Called(text, channel)
	return args.Error(0)
}
