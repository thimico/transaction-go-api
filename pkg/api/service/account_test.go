package service

import (
	"context"
	"projeto-pismo-api-go/pkg/api/service/mocks"
	"reflect"
	"testing"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"projeto-pismo-api-go/pkg/api/model"
)

func TestAccount_Create(t *testing.T) {
	mockRepo := mocks.NewMockAccountRepository()
	accountService := NewAccount(mockRepo)

	ctx := context.Background()
	accountIn := &model.AccountIn{
		DocumentNumber: "123456789",
	}

	accountOut, err := accountService.Create(ctx, accountIn)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if accountOut.DocumentNumber != accountIn.DocumentNumber {
		t.Errorf("expected document number %v, got %v", accountIn.DocumentNumber, accountOut.DocumentNumber)
	}

	if accountOut.ID.IsZero() {
		t.Errorf("expected non-zero account ID")
	}
}

func TestAccount_GetByID(t *testing.T) {
	mockRepo := mocks.NewMockAccountRepository()
	accountService := NewAccount(mockRepo)

	ctx := context.Background()
	accountIn := &model.AccountIn{
		DocumentNumber: "123456789",
	}

	// Create an account first
	createdAccount, err := accountService.Create(ctx, accountIn)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	// Retrieve the account by ID
	retrievedAccount, err := accountService.GetByID(ctx, createdAccount.ID.Hex())
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if !reflect.DeepEqual(createdAccount, retrievedAccount) {
		t.Errorf("expected account %v, got %v", createdAccount, retrievedAccount)
	}
}

func TestAccount_GetByID_NotFound(t *testing.T) {
	mockRepo := mocks.NewMockAccountRepository()
	accountService := NewAccount(mockRepo)

	ctx := context.Background()
	nonExistentID := primitive.NewObjectID().Hex()

	_, err := accountService.GetByID(ctx, nonExistentID)
	if err == nil {
		t.Fatalf("expected error, got nil")
	}

	expectedError := "account not found"
	if err.Error() != expectedError {
		t.Errorf("expected error %v, got %v", expectedError, err.Error())
	}
}
