package assistant

import (
	"cloud.google.com/go/functions/metadata"
	"context"
	"fmt"
	"github.com/jmrosh/go-genai-slack-app/api/services"
	firestoreModels "github.com/jmrosh/go-genai-slack-app/models/firestore"
	"log"
	"strings"
)

const (
	notFoundError = "not-found"
)

type Assistant interface {
	Run(ctx context.Context, event firestoreModels.EventDto) error
}

type assistant struct {
	FirestoreService services.FirestoreService
	SlackService     services.SlackService
	OpenAIService    services.OpenAIService
}

func NewAssistant(firestoreService services.FirestoreService, slackService services.SlackService, openAIService services.OpenAIService) Assistant {
	return &assistant{
		FirestoreService: firestoreService,
		SlackService:     slackService,
		OpenAIService:    openAIService,
	}
}

// Run is triggered by a change to a Firestore collection.
func (a *assistant) Run(ctx context.Context, event firestoreModels.EventDto) error {
	meta, err := metadata.FromContext(ctx)
	if err != nil {
		return fmt.Errorf("metadata.FromContext: %v", err)
	}

	//TODO: handle different event types
	//meta.EventType = "google.firestore.document.write"
	log.Printf("Function triggered by change to: %v\n", *meta.Resource)

	lastIndex := len(event.Value.Fields.Messages.ArrayValue.Values) - 1

	roleText := event.Value.Fields.Messages.ArrayValue.Values[lastIndex].MapValue.Fields.Role.StringValue

	if roleText == "Assistant" {
		log.Printf("Ignored Assistant message\n")
		return nil
	}

	messageText := event.Value.Fields.Messages.ArrayValue.Values[lastIndex].MapValue.Fields.Text.StringValue

	response, err := a.OpenAIService.GetChatCompletion(messageText)
	if response == nil || err != nil {
		return fmt.Errorf("error in GetChatCompletion(): %v", err)
	}

	path := strings.Split(meta.Resource.RawPath, "/")
	docId := path[len(path)-1]

	err = a.SlackService.SendMessage(*response, docId)
	if err != nil {
		return fmt.Errorf("error in SendMessage() for channelId: %v: %v", docId, err)
	}

	err = a.FirestoreService.AddMessage(ctx, "conversations", docId, "Assistant", *response)
	if err != nil {
		return fmt.Errorf("error in AddMessage(): %v", err)
	}

	return nil
}
