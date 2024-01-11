package run

import (
	"cloud.google.com/go/firestore"
	"context"
	"fmt"
	"github.com/jmrosh/go-genai-slack-app/api/services"
	"github.com/jmrosh/go-genai-slack-app/functions/assistant"
	"github.com/jmrosh/go-genai-slack-app/functions/subscriber"
	firestoreModels "github.com/jmrosh/go-genai-slack-app/models/firestore"
	"github.com/sashabaranov/go-openai"
	"github.com/slack-go/slack"
	"log"
	"net/http"
	"os"
)

func Assistant(ctx context.Context, e firestoreModels.EventDto) error {
	openAIClient := openai.NewClient(os.Getenv("OPENAI_API_KEY"))
	openAIService := services.NewOpenAIService(openAIClient)

	slackToken := os.Getenv("SLACK_AUTH_TOKEN")
	slackClient := slack.New(slackToken, slack.OptionDebug(true))
	slackService := services.NewSlackService(slackClient)

	firestoreClient, err := firestore.NewClient(ctx, os.Getenv("PROJECT_ID"))
	if err != nil {
		return fmt.Errorf("services.NewFirestoreClient: %v\n", err)
	}

	firestoreService := services.NewFirestoreService(ctx, firestoreClient)

	function := assistant.NewAssistant(firestoreService, slackService, openAIService)
	if err := function.Run(ctx, e); err != nil {
		return fmt.Errorf("assistant.Run: %v\n", err)
	}

	return nil
}

func Subscriber(w http.ResponseWriter, r *http.Request) {
	firestoreClient, err := firestore.NewClient(r.Context(), os.Getenv("PROJECT_ID"))
	if err != nil {
		log.Printf("services.NewFirestoreClient: %v\n", err)
		return
	}

	firestoreService := services.NewFirestoreService(context.Background(), firestoreClient)

	subscriber := subscriber.NewSubscriber(firestoreService)
	subscriber.Run(r.Context(), w, r)
}
