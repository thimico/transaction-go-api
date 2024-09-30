package controller

import (
	"context"
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"net/http"
	"projeto-pismo-api-go/pkg/api/model"
	"projeto-pismo-api-go/pkg/api/service"
)

type TransactionCtrl interface {
	CreateTransaction(w http.ResponseWriter, r *http.Request)
}

type TransactionCtrlImpl struct {
	service service.Transaction
}

func NewTransactionController(service service.Transaction) *TransactionCtrlImpl {
	return &TransactionCtrlImpl{service: service}
}

func (p *TransactionCtrlImpl) CreateTransaction(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var in model.TransactionIn
	err := json.NewDecoder(r.Body).Decode(&in)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	error := ValidateTransaction(&in)
	if error != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	transaction, err := p.service.Create(context.Background(), &in)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)

	encoder := json.NewEncoder(w)
	encoder.Encode(transaction)
}

func ValidateTransaction(v interface{}) error {
	var validate *validator.Validate
	validate = validator.New()

	errs := validate.Struct(v)
	if errs != nil {
		return errs
	}
	return nil
}
