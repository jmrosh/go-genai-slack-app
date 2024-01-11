package subscriber

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/jmrosh/go-genai-slack-app/api/services"
	"github.com/jmrosh/go-genai-slack-app/models/slack"
	"log"
	"net/http"
)

type Subscriber interface {
	Run(ctx context.Context, w http.ResponseWriter, r *http.Request)
}

type subscriber struct {
	FirestoreService services.FirestoreService
}

func NewSubscriber(firestoreService services.FirestoreService) Subscriber {
	return &subscriber{
		FirestoreService: firestoreService,
	}
}

func (s *subscriber) Run(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	e := slack.Event{}
	if err := json.NewDecoder(r.Body).Decode(&e); err != nil {
		log.Printf("error in json.Decode(): %v", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	if e.Type == "url_verification" {
		fmt.Fprint(w, e.Challenge)
		return
	}

	if e.InnerEvent.BotId != "" {
		fmt.Fprint(w, "Ignored bot event")
		return
	}

	err := s.FirestoreService.AddMessage(ctx, "conversations", e.InnerEvent.Channel, "User", e.InnerEvent.Text)
	if err != nil {
		log.Printf("error in AddMessage(): %v\n", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
}
