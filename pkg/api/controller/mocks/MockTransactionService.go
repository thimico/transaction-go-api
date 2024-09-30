package mocks

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
