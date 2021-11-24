package account_service

import (
	"errors"
	"pismo-desafio-app/app/model/account"
	"pismo-desafio-app/app/repository/account_repository"
)

func Create(a account.Account) (account.Account, error) {
	accountDb, _ := account_repository.GetByDocumentNumber(a.DocNumber)
	if len(accountDb) > 0 {
		return account.Account{}, errors.New("conta já cadastrada")
	}
	return account_repository.Create(a)
}

func GetById(id int) (account.Account, error) {
	return account_repository.GetById(id)
}
