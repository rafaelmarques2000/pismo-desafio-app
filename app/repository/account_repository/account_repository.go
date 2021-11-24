package account_repository

import (
	"pismo-desafio-app/app/model/account"
	"pismo-desafio-app/config/database"
)

func Create(a account.Account) (account.Account, error) {
	connection, err := database.OpenConnection()
	if err != nil {
		return account.Account{}, err
	}
	result := connection.Create(&a)
	return a, result.Error
}

func GetById(id int) (account.Account, error) {
	connection, err := database.OpenConnection()
	if err != nil {
		return account.Account{}, err
	}
	var accountSearch account.Account
	result := connection.Model(&account.Account{}).First(&accountSearch, id)
	return accountSearch, result.Error
}

func GetByDocumentNumber(document string) ([]account.Account, error) {
	connection, err := database.OpenConnection()
	if err != nil {
		return []account.Account{}, err
	}
	var accountSearch []account.Account
	result := connection.Model(&account.Account{}).Where("doc_number = ?", document).Find(&accountSearch)
	return accountSearch, result.Error
}
