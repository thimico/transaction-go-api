package service

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"projeto-pismo-api-go/pkg/api/service/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"projeto-pismo-api-go/pkg/api/model"
)

func TestCreateTransactionService(t *testing.T) {
	mockRepo := &mocks.MockTransactionRepository{}
	service := NewTransaction(mockRepo)

	transactionIn := &model.TransactionIn{
		AccountID: primitive.NewObjectID(),
		Amount:    100,
	}

	transactionOut, err := service.Create(context.Background(), transactionIn)
	assert.NoError(t, err)
	assert.Equal(t, transactionIn.Amount, transactionOut.Amount)
}
