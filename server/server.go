package server

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"pismo-desafio-app/controllers/account"
	"pismo-desafio-app/controllers/transaction"
)

var router *mux.Router

func Start() {
	router = mux.NewRouter()
	router.Use(globalMiddleware)

	/** Accounts Routes **/
	router.HandleFunc("/accounts", account.CreateAccount).Methods("POST")
	router.HandleFunc("/accounts/{id}", account.GetAccountById).Methods("GET")

	/** Transactions Routes **/
	router.HandleFunc("/transactions", transaction.CreateTransaction).Methods("POST")

	log.Fatal(http.ListenAndServe(":8080", router))
}

func globalMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}