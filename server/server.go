package server

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"pismo-desafio-app/pkg/controller"
)

type Server struct {
	accountController *controller.AccountController
	transactionController *controller.TransactionController
	router *mux.Router
}

func NewServer() *Server {
	return &Server{
		accountController:     controller.NewAccountController(),
		transactionController: controller.NewTransactionController(),
		router:                mux.NewRouter(),
	}
}

func (s *Server)Start() {
	s.router.Use(globalMiddleware)

	//Accounts Routes
	s.router.HandleFunc("/accounts", s.accountController.CreateAccount).Methods("POST")
	s.router.HandleFunc("/accounts/{id}", s.accountController.GetAccountById).Methods("GET")

	//Transactions Routes
	s.router.HandleFunc("/transactions", s.transactionController.CreateTransaction).Methods("POST")
	log.Fatal(http.ListenAndServe(os.Getenv("SERVER_PORT"), s.router))
}

func globalMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
