package mocks

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"projeto-pismo-api-go/pkg/api/model"
)

type MockAccountService struct{}

func (m *MockAccountService) Create(ctx context.Context, in *model.AccountIn) (*model.AccountOut, error) {
	return &model.AccountOut{
		ID:             primitive.NewObjectID(),
		DocumentNumber: in.DocumentNumber,
	}, nil
}

func (m *MockAccountService) GetByID(ctx context.Context, id string) (*model.AccountOut, error) {
	return &model.AccountOut{
		ID:             primitive.NewObjectID(),
		DocumentNumber: "12345678900",
	}, nil
}
