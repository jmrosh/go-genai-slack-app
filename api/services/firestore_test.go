package services

import (
	"cloud.google.com/go/firestore"
	"context"
	"github.com/jmrosh/go-genai-slack-app/api/services/mocks"
	"github.com/stretchr/testify/mock"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFirestoreService_AddMessage(t *testing.T) {
	ctx := context.Background()

	var defaultMockClient = mocks.NewFirestoreClient()
	defaultService := NewFirestoreService(ctx, defaultMockClient)

	t.Run("Correct inputs", func(t *testing.T) {
		// Arrange
		var mockClient = mocks.NewFirestoreClient()
		service := NewFirestoreService(ctx, mockClient)

		mockClient.On("Collection", mock.AnythingOfType("string")).Return(&firestore.CollectionRef{
			ID: "123",
		})
		mockClient.On("Doc", mock.AnythingOfType("string")).Return(&firestore.DocumentRef{
			ID: "123",
		})
		mockClient.On("RunTransaction", mock.Anything, mock.Anything, mock.Anything).Return(nil)

		// Act
		err := service.AddMessage(ctx, "collection1", "document1", "role1", "text1")

		// Assert
		assert.NoError(t, err)
	})

	t.Run("Empty collection name", func(t *testing.T) {
		// Arrange
		var mockClient = mocks.NewFirestoreClient()
		service := NewFirestoreService(ctx, mockClient)
		mockClient.On("Collection", "").Return(&firestore.CollectionRef{})

		// Act
		err := service.AddMessage(ctx, "", "document2", "role2", "text2")

		// Assert
		assert.Error(t, err)
	})

	t.Run("Empty document ID", func(t *testing.T) {
		// Arrange
		var mockClient = mocks.NewFirestoreClient()
		service := NewFirestoreService(ctx, mockClient)

		mockClient.On("Collection", mock.AnythingOfType("string")).Return(&firestore.CollectionRef{
			ID: "123",
		})
		mockClient.On("doc", "").Return(&firestore.DocumentRef{})

		// Act
		err := service.AddMessage(ctx, "collection3", "", "role3", "text3")

		// Assert
		assert.Error(t, err)
	})

	t.Run("Empty role", func(t *testing.T) {
		// Act
		err := defaultService.AddMessage(ctx, "collection4", "document4", "", "text4")

		// Assert
		assert.Error(t, err)
	})

	t.Run("Empty text", func(t *testing.T) {
		// Act
		err := defaultService.AddMessage(ctx, "collection5", "document5", "role5", "")

		// Assert
		assert.Error(t, err)
	})
}
