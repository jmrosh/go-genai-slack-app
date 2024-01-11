package mocks

import (
	"cloud.google.com/go/firestore"
	"context"
	"github.com/stretchr/testify/mock"
)

type MockFirestoreClient struct {
	mock.Mock
}

type MockFirestoreService struct {
	mock.Mock
}

func NewFirestoreClient() *MockFirestoreClient {
	return &MockFirestoreClient{}
}

func NewFirestoreService() *MockFirestoreService {
	return &MockFirestoreService{}
}

func (m *MockFirestoreClient) RunTransaction(ctx context.Context, f func(context.Context, *firestore.Transaction) error, opts ...firestore.TransactionOption) error {
	args := m.Called(ctx, f, opts)
	return args.Error(0)
}

func (m *MockFirestoreClient) Collection(path string) *firestore.CollectionRef {
	args := m.Called(path)
	return args.Get(0).(*firestore.CollectionRef)
}

func (m *MockFirestoreClient) Doc(path string) *firestore.DocumentRef {
	args := m.Called(path)
	return args.Get(0).(*firestore.DocumentRef)
}

func (m *MockFirestoreService) AddMessage(ctx context.Context, collectionId string, documentId string, role string, text string) error {
	args := m.Called(ctx, collectionId, documentId, role, text)
	if args.Error(0) != nil {
		return args.Error(0)
	}
	return nil
}
