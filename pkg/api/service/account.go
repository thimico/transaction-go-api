package service

import (
	"context"
	"projeto-pismo-api-go/pkg/api/model"
)

type Account interface {
	Create(ctx context.Context, account *model.AccountIn) (*model.AccountOut, error)
	GetByID(ctx context.Context, id string) (*model.AccountOut, error)
}

type AccountRepository interface {
	Create(ctx context.Context, account *model.Account) (*model.Account, error)
	GetByID(ctx context.Context, id string) (*model.Account, error)
}

type account struct {
	repository AccountRepository
}

func NewAccount(repository AccountRepository) Account {
	return &account{repository: repository}
}

func (s *account) Create(ctx context.Context, in *model.AccountIn) (*model.AccountOut, error) {
	accountModel := in.ToAccount()
	account, err := s.repository.Create(ctx, accountModel)
	if err != nil {
		return nil, err
	}
	return account.ToAccountOut(), nil
}

func (s *account) GetByID(ctx context.Context, id string) (*model.AccountOut, error) {
	account, err := s.repository.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return account.ToAccountOut(), nil
}