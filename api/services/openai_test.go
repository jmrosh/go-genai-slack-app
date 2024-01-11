package services

import (
	"github.com/jmrosh/go-genai-slack-app/api/services/mocks"
	"github.com/sashabaranov/go-openai"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestOpenAIService_GetChatCompletion(t *testing.T) {
	mockClient := mocks.NewMockOpenAIClient()
	service := NewOpenAIService(mockClient)

	t.Run("GetChatCompletion_ExpectedReturnValue", func(t *testing.T) {
		// Arrange
		expectedResponse := "response"
		mockClient.On("CreateChatCompletion", mock.Anything, mock.Anything).Return(openai.ChatCompletionResponse{
			Choices: []openai.ChatCompletionChoice{
				{
					Message: openai.ChatCompletionMessage{
						Content: expectedResponse,
					},
				},
			},
		}, nil)

		// Act
		output, err := service.GetChatCompletion("message")

		// Assert
		assert.Equal(t, expectedResponse, *output)
		assert.NoError(t, err)
		assert.NotNil(t, output, "Unexpected nil output")
	})

	t.Run("GetChatCompletion_WithEmptyMessage", func(t *testing.T) {
		// Act
		output, err := service.GetChatCompletion("")

		// Assert
		assert.Nil(t, output, "Unexpected non-nil output")
		assert.Error(t, err)
	})
}
