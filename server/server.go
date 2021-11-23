package server

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"pismo-desafio-app/controllers/account"
	"pismo-desafio-app/controllers/default"
	"pismo-desafio-app/controllers/transaction"
)

var router *mux.Router

func Start() {
	router = mux.NewRouter()

	/** Default Routes **/
	router.HandleFunc("/", _default.DefaultIndex).Methods("GET")

	/** Accounts Routes **/
	router.HandleFunc("/accounts", account.CreateAccount).Methods("POST")
	router.HandleFunc("/accounts", account.GetAccountById).Methods("GET")

	/** Transactions Routes **/
	router.HandleFunc("/transactions", transaction.CreateTransaction).Methods("POST")

	log.Fatal(http.ListenAndServe(":8080", router))
}