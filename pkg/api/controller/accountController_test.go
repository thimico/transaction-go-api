package controller

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"projeto-pismo-api-go/pkg/api/controller/mocks"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"projeto-pismo-api-go/pkg/api/model"
)

func TestCreateAccount(t *testing.T) {
	mockService := &mocks.MockAccountService{}
	controller := NewAccountController(mockService)

	accountIn := &model.AccountIn{
		DocumentNumber: "12345678900",
	}
	body, _ := json.Marshal(accountIn)
	req, err := http.NewRequest("POST", "/accounts", bytes.NewBuffer(body))
	assert.NoError(t, err)
	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(controller.CreateAccount)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusCreated, rr.Code)
	var accountOut model.AccountOut
	err = json.NewDecoder(rr.Body).Decode(&accountOut)
	assert.NoError(t, err)
	assert.Equal(t, accountIn.DocumentNumber, accountOut.DocumentNumber)
}

func TestGetAccount(t *testing.T) {
	mockService := &mocks.MockAccountService{}
	controller := NewAccountController(mockService)

	req, err := http.NewRequest("GET", "/accounts/{id}", nil)
	assert.NoError(t, err)
	rr := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/accounts/{id}", controller.GetAccount)
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	var accountOut model.AccountOut
	err = json.NewDecoder(rr.Body).Decode(&accountOut)
	assert.NoError(t, err)
	assert.Equal(t, "12345678900", accountOut.DocumentNumber)
}
