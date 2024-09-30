package mocks

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"projeto-pismo-api-go/pkg/api/model"
)

// MockAccountRepository is a mock implementation of the AccountRepository interface
type MockAccountRepository struct {
	accounts map[string]*model.Account
}

func NewMockAccountRepository() *MockAccountRepository {
	return &MockAccountRepository{
		accounts: make(map[string]*model.Account),
	}
}

func (m *MockAccountRepository) Create(ctx context.Context, account *model.Account) (*model.Account, error) {
	account.ID = primitive.NewObjectID()
	m.accounts[account.ID.Hex()] = account
	return account, nil
}

func (m *MockAccountRepository) GetByID(ctx context.Context, id string) (*model.Account, error) {
	account, exists := m.accounts[id]
	if !exists {
		return nil, errors.New("account not found")
	}
	return account, nil
}
