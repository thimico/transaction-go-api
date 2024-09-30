package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Account struct {
	ID             primitive.ObjectID `bson:"_id,omitempty" json:"account_id"`
	DocumentNumber string             `bson:"document_number" json:"document_number"`
}

type AccountIn struct {
	DocumentNumber string `json:"document_number" validate:"required"`
}

type AccountOut struct {
	ID             primitive.ObjectID `json:"account_id"`
	DocumentNumber string             `json:"document_number"`
}

func (a *AccountIn) ToAccount() *Account {
	return &Account{
		DocumentNumber: a.DocumentNumber,
	}
}

func (a *Account) ToAccountOut() *AccountOut {
	return &AccountOut{
		ID:             a.ID,
		DocumentNumber: a.DocumentNumber,
	}
}