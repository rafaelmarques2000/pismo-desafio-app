package transaction_service

import (
	"pismo-desafio-app/app/model/account"
	"pismo-desafio-app/app/model/operationtype"
	"pismo-desafio-app/app/model/transaction"
	"pismo-desafio-app/app/repository/transaction_repository"
	"pismo-desafio-app/app/service/account_service"
	"pismo-desafio-app/app/service/operation_type_service"
	"time"
)

func Create(t transaction.ApiTransaction) (transaction.ApiTransaction, error) {
	accountDb, err := account_service.GetById(t.AccountId)
	if err != nil {
		return transaction.ApiTransaction{}, err
	}

	OperationType, err := operation_type_service.GetById(t.OperationTypeId)
	if err != nil {
		return transaction.ApiTransaction{}, err
	}

	return createTransactionDb(t, accountDb, OperationType, err)
}

func createTransactionDb(t transaction.ApiTransaction, accountDb account.Account, OperationType operationtype.OperationType, err error) (transaction.ApiTransaction, error) {
	var transactionDb = transaction.Transaction{
		Account:       accountDb,
		OperationType: OperationType,
		Amount:        t.Amount,
		EventDate:     time.Now(),
	}

	transactionError := transactionDb.CheckTransaction()
	if transactionError != nil {
		return transaction.ApiTransaction{}, transactionError
	}

	tx, err := transaction_repository.Create(transactionDb)
	if err != nil {
		return transaction.ApiTransaction{}, err
	}

	t.EventDate = tx.EventDate
	t.Id = tx.ID

	return t, err
}


