package services

import (
	"cloud.google.com/go/firestore"
	"context"
	"fmt"
	firestoreModels "github.com/jmrosh/go-genai-slack-app/models/firestore"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
)

type FirestoreService interface {
	AddMessage(ctx context.Context, collectionId string, documentId string, role string, text string) error
}

type FirestoreClient interface {
	RunTransaction(ctx context.Context, f func(context.Context, *firestore.Transaction) error, opts ...firestore.TransactionOption) (err error)
	Collection(path string) *firestore.CollectionRef
	Doc(path string) *firestore.DocumentRef
}

type firestoreService struct {
	Context context.Context
	Client  FirestoreClient
}

func NewFirestoreService(ctx context.Context, client FirestoreClient) FirestoreService {
	return &firestoreService{
		Context: ctx,
		Client:  client,
	}
}

const (
	notFoundError = "not-found"
)

func (s *firestoreService) AddMessage(ctx context.Context, collectionId string, documentId string, role string, text string) error {
	if text == "" {
		return fmt.Errorf("text cannot be empty")
	}
	if role == "" {
		return fmt.Errorf("role cannot be empty")
	}

	collection := s.Client.Collection(collectionId)
	if collection == nil || collection.ID == "" {
		return fmt.Errorf("error in Collection(): no collection found")
	}

	docRef := collection.Doc(documentId)
	if docRef == nil || docRef.ID == "" {
		return fmt.Errorf("error in Doc(): no doc found")
	}

	err := s.Client.RunTransaction(ctx, func(ctx context.Context, tx *firestore.Transaction) error {
		doc, err := tx.Get(docRef)
		if err != nil && status.Code(err) != codes.NotFound {
			return err
		}

		if doc.Exists() {
			// Update the document
			message := firestoreModels.Message{Role: role, Text: text}
			return tx.Update(docRef, []firestore.Update{
				{Path: "Messages", Value: firestore.ArrayUnion(message)},
			})
		}

		// Else create the document
		return tx.Set(docRef, firestoreModels.Conversation{
			Messages: []firestoreModels.Message{{Role: role, Text: text}},
		})
	})

	if err != nil {
		return fmt.Errorf("error in RunTransaction(): %v", err)
	}

	log.Printf("Added or updated document with ID: %v\n", docRef.ID)
	return nil
}
