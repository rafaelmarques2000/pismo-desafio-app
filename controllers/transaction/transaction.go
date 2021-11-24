package transaction

import (
	"encoding/json"
	"net/http"
	"pismo-desafio-app/app/model/transaction"
	"pismo-desafio-app/app/service/transaction_service"
	"pismo-desafio-app/util/response"
)

func CreateTransaction(w http.ResponseWriter, r *http.Request) {

	var transactionRequest transaction.ApiTransaction
	err := json.NewDecoder(r.Body).Decode(&transactionRequest)

	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode(response.Response{err.Error()})
		return
	}

	transactionDb, err := transaction_service.Create(transactionRequest)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode(response.Response{Msg: err.Error()})
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response.ResponseSuccess{
		"Transação realizada com sucesso",
		transactionDb,
	})
}
