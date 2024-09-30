package controller

import (
	"bytes"
	"context"
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"projeto-pismo-api-go/pkg/api/model"
)

type MockTransactionService struct{}

func (m *MockTransactionService) Create(ctx context.Context, in *model.TransactionIn) (*model.TransactionOut, error) {
	return &model.TransactionOut{
		ID:        primitive.NewObjectID(),
		AccountID: in.AccountID,
		Amount:    in.Amount,
	}, nil
}

func TestCreateTransaction(t *testing.T) {
	mockService := &MockTransactionService{}
	controller := NewTransactionController(mockService)

	transactionIn := &model.TransactionIn{
		AccountID: primitive.NewObjectID(),
		Amount:    100,
	}
	body, _ := json.Marshal(transactionIn)
	req, err := http.NewRequest("POST", "/transactions", bytes.NewBuffer(body))
	assert.NoError(t, err)
	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(controller.CreateTransaction)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusCreated, rr.Code)
	var transactionOut model.TransactionOut
	err = json.NewDecoder(rr.Body).Decode(&transactionOut)
	assert.NoError(t, err)
	assert.Equal(t, transactionIn.Amount, transactionOut.Amount)
}
