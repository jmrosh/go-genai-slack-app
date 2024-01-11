package subscriber

import (
	"bytes"
	"context"
	"errors"
	"github.com/jmrosh/go-genai-slack-app/api/services/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSubscriber(t *testing.T) {
	t.Run("When url verification responds with challenge and returns 200", func(t *testing.T) {
		service := mocks.NewFirestoreService()
		s := NewSubscriber(service)

		req, _ := http.NewRequest(http.MethodPost, "/", bytes.NewBufferString(`{"type": "url_verification", "challenge": "sampleChallenge"}`))
		rr := httptest.NewRecorder()

		s.Run(context.Background(), rr, req)

		assert.Equal(t, http.StatusOK, rr.Code)
		assert.Equal(t, "sampleChallenge", rr.Body.String())
	})

	t.Run("When bot event ignores and returns 200", func(t *testing.T) {
		service := mocks.NewFirestoreService()
		s := NewSubscriber(service)

		req, _ := http.NewRequest(http.MethodPost, "/", bytes.NewBufferString(`{"type": "message", "event": { "bot_id": "bot1"}}`))
		rr := httptest.NewRecorder()

		s.Run(context.Background(), rr, req)

		assert.Equal(t, http.StatusOK, rr.Code)
		assert.Equal(t, "Ignored bot event", rr.Body.String())
	})

	t.Run("When valid slack message event returns 200", func(t *testing.T) {
		service := mocks.NewFirestoreService()
		service.On("AddMessage", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil)

		s := NewSubscriber(service)

		req, _ := http.NewRequest(http.MethodPost, "/", bytes.NewBufferString(`{"type": "message", "event": {"text": "hello", "channel": "channel1"}}`))
		rr := httptest.NewRecorder()

		s.Run(context.Background(), rr, req)

		assert.Equal(t, http.StatusOK, rr.Code)
	})

	t.Run("When firestore fails returns 500", func(t *testing.T) {
		service := mocks.NewFirestoreService()
		service.On("AddMessage", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(errors.New("unknown error"))

		s := NewSubscriber(service)

		req, _ := http.NewRequest(http.MethodPost, "/", bytes.NewBufferString(`{"type": "message", "event": {"text": "hello", "channel": "channel1"}}`))
		rr := httptest.NewRecorder()

		s.Run(context.Background(), rr, req)

		assert.Equal(t, http.StatusInternalServerError, rr.Code)
	})
}
