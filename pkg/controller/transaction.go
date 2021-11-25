package controller

import (
	"encoding/json"
	"net/http"
	"pismo-desafio-app/pkg/model"
	"pismo-desafio-app/pkg/service"
	"pismo-desafio-app/pkg/util"
)

type TransactionController struct {
	TransactionService *service.TransactionService
}

func NewTransactionController() *TransactionController {
	return &TransactionController{
		TransactionService: service.NewTransactionService(),
	}
}

func (t *TransactionController) CreateTransaction(w http.ResponseWriter, r *http.Request) {

	var transactionRequest model.ApiTransaction
	err := json.NewDecoder(r.Body).Decode(&transactionRequest)

	if err != nil {
		util.NewApiResponse(w, http.StatusUnprocessableEntity, util.Response{Msg: err.Error()})
		return
	}

	transactionDb, err := t.TransactionService.Create(transactionRequest)
	if err != nil {
		util.NewApiResponse(w, http.StatusUnprocessableEntity, util.Response{Msg: err.Error()})
		return
	}

	util.NewApiResponse(w, http.StatusCreated, util.ResponseSuccess{
		Msg:  "Transação realizada com sucesso",
		Data: transactionDb.GetApiTransaction(),
	})
}
