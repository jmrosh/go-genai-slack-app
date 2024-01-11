package services

import (
	"context"
	"fmt"
	openai "github.com/sashabaranov/go-openai"
)

type OpenAIService interface {
	GetChatCompletion(message string) (*string, error)
}

type OpenAIClient interface {
	CreateChatCompletion(ctx context.Context, request openai.ChatCompletionRequest) (response openai.ChatCompletionResponse, err error)
}

type openAIService struct {
	client OpenAIClient
}

func NewOpenAIService(client OpenAIClient) OpenAIService {
	return &openAIService{
		client: client,
	}
}

// GetChatCompletion sends a request to the ChatCompletion API and returns the first response choice.
func (s *openAIService) GetChatCompletion(message string) (*string, error) {
	if message == "" {
		return nil, fmt.Errorf("message cannot be empty\n")
	}

	request := getPrompt(message)

	// make request
	ctx := context.Background()
	resp, err := s.client.CreateChatCompletion(ctx, request)
	if err != nil {
		return nil, fmt.Errorf("error in client.CreateChatCompletion: %v\n", err)
	}

	// Return first response
	return &resp.Choices[0].Message.Content, nil
}

func getPrompt(message string) openai.ChatCompletionRequest {
	return openai.ChatCompletionRequest{
		Model:       "gpt-4-1106-preview",
		N:           1,
		Temperature: .80,
		Messages: []openai.ChatCompletionMessage{
			{
				Role: "system",
				Content: "You are an assistant running in a Slack app. So take advantage of Slack's formatting" +
					" conventions to produce rich responses.",
			},
			{
				Role: "system",
				Content: "Your user is a software engineer at Slalom Build. The simpler the question you receive, " +
					"the more concise your response is.",
			},
			{
				Role:    "user",
				Content: message,
			},
		},
	}
}
