package subscriber

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/jmrosh/go-genai-slack-app/api/services"
	"github.com/slack-go/slack/slackevents"
	"io"
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
	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	event, err := slackevents.ParseEvent(json.RawMessage(body), slackevents.OptionNoVerifyToken())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&event); err != nil {
		log.Printf("error in json.Decode(): %v", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	if event.Type == slackevents.URLVerification {
		var r *slackevents.ChallengeResponse
		err := json.Unmarshal(body, &r)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "text")
		w.Write([]byte(r.Challenge))
		return
	}

	switch innerEvent := event.InnerEvent.Data.(type) {
	case *slackevents.MessageEvent:
		if innerEvent.BotID != "" {
			fmt.Fprint(w, "Ignored bot event")
			return
		}
		err := s.FirestoreService.AddMessage(ctx, "conversations", innerEvent.Channel, "User", innerEvent.Text)
		if err != nil {
			log.Printf("error in AddMessage(): %v\n", err)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
	}
}
