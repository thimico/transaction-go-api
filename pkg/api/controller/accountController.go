package controller

import (
	"context"
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"net/http"
	"projeto-pismo-api-go/pkg/api/model"
	"projeto-pismo-api-go/pkg/api/service"
)

type AccountCtrl interface {
	CreateAccount(w http.ResponseWriter, r *http.Request)
	GetAccount(w http.ResponseWriter, r *http.Request)
}

type AccountCtrlImpl struct {
	service service.Account
}

func NewAccountController(service service.Account) *AccountCtrlImpl {
	return &AccountCtrlImpl{service: service}
}

func (p *AccountCtrlImpl) CreateAccount(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var in model.AccountIn
	err := json.NewDecoder(r.Body).Decode(&in)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	error := ValidateAccount(&in)
	if error != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}
	account, err := p.service.Create(context.Background(), &in)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	encoder := json.NewEncoder(w)
	encoder.Encode(account)
	w.WriteHeader(http.StatusCreated)
}

func (p *AccountCtrlImpl) GetAccount(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	account, err := p.service.GetByID(context.Background(), vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(account)
}

func ValidateAccount(v interface{}) error {
	var validate *validator.Validate
	validate = validator.New()

	errs := validate.Struct(v)
	if errs != nil {
		return errs
	}
	return nil
}
