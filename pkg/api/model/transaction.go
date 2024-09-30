package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Transaction struct {
	ID             primitive.ObjectID `bson:"_id,omitempty" json:"transaction_id"`
	AccountID      primitive.ObjectID `bson:"account_id" json:"account_id"`
	OperationTypeID int               `bson:"operation_type_id" json:"operation_type_id"`
	Amount         float64            `bson:"amount" json:"amount"`
	EventDate      time.Time          `bson:"event_date" json:"event_date"`
}

type TransactionIn struct {
	AccountID      primitive.ObjectID `json:"account_id" validate:"required"`
	OperationTypeID int               `json:"operation_type_id" validate:"required"`
	Amount         float64            `json:"amount" validate:"required"`
}

type TransactionOut struct {
	ID             primitive.ObjectID `json:"transaction_id"`
	AccountID      primitive.ObjectID `json:"account_id"`
	OperationTypeID int               `json:"operation_type_id"`
	Amount         float64            `json:"amount"`
	EventDate      time.Time          `json:"event_date"`
}

func (t *TransactionIn) ToTransaction() *Transaction {
	return &Transaction{
		AccountID:      t.AccountID,
		OperationTypeID: t.OperationTypeID,
		Amount:         t.Amount,
		EventDate:      time.Now(),
	}
}

func (t *Transaction) ToTransactionOut() *TransactionOut {
	return &TransactionOut{
		ID:             t.ID,
		AccountID:      t.AccountID,
		OperationTypeID: t.OperationTypeID,
		Amount:         t.Amount,
		EventDate:      t.EventDate,
	}
}