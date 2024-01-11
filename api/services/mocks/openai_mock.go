package mocks

import (
	"context"
	"github.com/sashabaranov/go-openai"
	"github.com/stretchr/testify/mock"
)

type MockOpenAIClient struct {
	mock.Mock
}

type MockOpenAIService struct {
	mock.Mock
}

func NewMockOpenAIClient() *MockOpenAIClient {
	return &MockOpenAIClient{}
}

func NewMockOpenAIService() *MockOpenAIService {
	return &MockOpenAIService{}
}

func (m *MockOpenAIClient) CreateChatCompletion(ctx context.Context, request openai.ChatCompletionRequest) (response openai.ChatCompletionResponse, err error) {
	args := m.Called(ctx, request)
	return args.Get(0).(openai.ChatCompletionResponse), args.Error(1)
}

func (m *MockOpenAIService) GetChatCompletion(message string) (*string, error) {
	args := m.Called(message)
	return args.Get(0).(*string), args.Error(1)
}
