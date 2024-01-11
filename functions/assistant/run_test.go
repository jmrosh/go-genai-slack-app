package assistant

import (
	"cloud.google.com/go/functions/metadata"
	"context"
	"errors"
	"github.com/jmrosh/go-genai-slack-app/api/services/mocks"
	firestoreModels "github.com/jmrosh/go-genai-slack-app/models/firestore"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestAssistant_Run(t *testing.T) {
	md := metadata.Metadata{
		Resource: &metadata.Resource{
			Name: "projects/test-project/databases/(default)/documents/conversations/channel1",
		},
		EventID:   "123",
		EventType: "google.firestore.document.write",
	}

	ctx := metadata.NewContext(context.Background(), &md)
	var nilStr *string = nil
	message := "Hello"
	event := firestoreModels.EventDto{
		Value: firestoreModels.Value{
			Name: "projects/test-project/databases/(default)/documents/conversations/channel1",
			Fields: firestoreModels.ConversationDto{
				Messages: firestoreModels.MessageContainer{
					ArrayValue: firestoreModels.MessageArrayContainer{
						Values: []firestoreModels.MessageValue{
							{
								MapValue: firestoreModels.MessageMapValue{
									Fields: firestoreModels.MessageFields{
										Role: firestoreModels.RoleValue{
											StringValue: "User",
										},
										Text: firestoreModels.TextValue{
											StringValue: message,
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}

	t.Run("When no service errors returns nil", func(t *testing.T) {
		openAiService := mocks.NewMockOpenAIService()
		firestoreService := mocks.NewFirestoreService()
		slackService := mocks.NewMockSlackService()
		assistant := NewAssistant(firestoreService, slackService, openAiService)

		slackService.On("SendMessage", mock.Anything, mock.Anything).Return(nil)
		openAiService.On("GetChatCompletion", mock.Anything).Return(&message, nil)
		firestoreService.On("AddMessage", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil)

		err := assistant.Run(ctx, event)
		assert.Nil(t, err)
	})

	t.Run("When open ai error returns error", func(t *testing.T) {
		openAiService := mocks.NewMockOpenAIService()
		firestoreService := mocks.NewFirestoreService()
		slackService := mocks.NewMockSlackService()
		assistant := NewAssistant(firestoreService, slackService, openAiService)
		openAiService.On("GetChatCompletion", mock.Anything).Return(nilStr, errors.New("error"))

		err := assistant.Run(ctx, event)
		assert.Error(t, err)
	})

	t.Run("When slack error returns error", func(t *testing.T) {
		openAiService := mocks.NewMockOpenAIService()
		firestoreService := mocks.NewFirestoreService()
		slackService := mocks.NewMockSlackService()
		assistant := NewAssistant(firestoreService, slackService, openAiService)

		openAiService.On("GetChatCompletion", mock.Anything).Return(nilStr, errors.New("error"))
		slackService.On("SendMessage", mock.Anything, mock.Anything).Return(errors.New("error"))

		err := assistant.Run(ctx, event)
		assert.Error(t, err)
	})

	t.Run("When firestore error returns error", func(t *testing.T) {
		openAiService := mocks.NewMockOpenAIService()
		firestoreService := mocks.NewFirestoreService()
		slackService := mocks.NewMockSlackService()
		assistant := NewAssistant(firestoreService, slackService, openAiService)

		slackService.On("SendMessage", mock.Anything, mock.Anything).Return(nil)
		openAiService.On("GetChatCompletion", mock.Anything).Return(&message, nil)
		firestoreService.On("AddMessage", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(errors.New("error"))

		err := assistant.Run(ctx, event)
		assert.Error(t, err)
	})

}
