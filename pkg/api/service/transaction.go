package service

import (
	"context"
	"projeto-pismo-api-go/pkg/api/model"
)

type Transaction interface {
	Create(ctx context.Context, transaction *model.TransactionIn) (*model.TransactionOut, error)
}

type TransactionRepository interface {
	Create(ctx context.Context, transaction *model.Transaction) (*model.Transaction, error)
}

type transaction struct {
	repository TransactionRepository
}

func NewTransaction(repository TransactionRepository) Transaction {
	return &transaction{repository: repository}
}

func (s *transaction) Create(ctx context.Context, in *model.TransactionIn) (*model.TransactionOut, error) {
	transactionModel := in.ToTransaction()
	transaction, err := s.repository.Create(ctx, transactionModel)
	if err != nil {
		return nil, err
	}
	return transaction.ToTransactionOut(), nil
}