package transaction_repository

import (
	"pismo-desafio-app/app/model/transaction"
	"pismo-desafio-app/config/database"
)

func Create(t transaction.Transaction) (transaction.Transaction, error) {
	connection, err := database.OpenConnection()
	if err != nil {
		return transaction.Transaction{}, err
	}
	result :=connection.Create(&t)
	return t, result.Error
}
