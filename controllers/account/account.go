package account

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"pismo-desafio-app/app/model/account"
	"pismo-desafio-app/app/service/account_service"
	"pismo-desafio-app/util/response"
	"strconv"
)

/** Criar uma nova conta **/
func CreateAccount(w http.ResponseWriter, r *http.Request) {
	var requestAccount account.Account
	err := json.NewDecoder(r.Body).Decode(&requestAccount)

	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode(response.Response{err.Error()})
		return
	}

	accountDb, err := account_service.Create(requestAccount)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode(response.Response{err.Error()})
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response.ResponseSuccess{
		"Conta criada com sucesso",
		accountDb,
	})
}

// Obter conta especifica pelo Id
func GetAccountById(w http.ResponseWriter, r *http.Request) {
	 var params map[string]string = mux.Vars(r)
	 id,_ := strconv.Atoi(params["id"])

	accountDb, err := account_service.GetById(id)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode(response.Response{err.Error()})
		return
	}

	json.NewEncoder(w).Encode(response.ResponseSuccess{
		"",
		accountDb,
	})
}