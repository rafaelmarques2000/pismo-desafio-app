package controller

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"pismo-desafio-app/pkg/model"
	"pismo-desafio-app/pkg/service"
	"pismo-desafio-app/pkg/util"
	"strconv"
)

type AccountController struct {
	AccountService *service.AccountService
}

func NewAccountController() *AccountController {
	return &AccountController{
		AccountService: service.NewAccountService(),
	}
}

/** Criar uma nova conta **/
func (a *AccountController) CreateAccount(w http.ResponseWriter, r *http.Request) {
	var requestAccount model.Account
	err := json.NewDecoder(r.Body).Decode(&requestAccount)

	if err != nil {
		util.NewApiResponse(w, http.StatusUnprocessableEntity, util.Response{Msg: err.Error()})
		return
	}

	accountDb, err := a.AccountService.Create(requestAccount)
	if err != nil {
		util.NewApiResponse(w, http.StatusUnprocessableEntity, util.Response{Msg: err.Error()})
		return
	}

	util.NewApiResponse(w, http.StatusCreated, util.ResponseSuccess{
		"Conta criada com sucesso",
		accountDb.GetApiAccount(),
	})
}

// Obter conta especifica pelo Id
func (a *AccountController) GetAccountById(w http.ResponseWriter, r *http.Request) {
	var params map[string]string = mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	accountDb, err := a.AccountService.GetById(uint(id))
	if err != nil {
		util.NewApiResponse(w, http.StatusUnprocessableEntity, util.Response{Msg: err.Error()})
		return
	}
	util.NewApiResponse(w, http.StatusOK, util.ResponseSuccess{
		Data: accountDb.GetApiAccount(),
	})
}
