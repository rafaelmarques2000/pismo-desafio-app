package service

import (
	"errors"
	"pismo-desafio-app/pkg/model"
	"pismo-desafio-app/pkg/repository"
	"time"
)

type TransactionService struct {
	AccountService        *AccountService
	OperationTypeService  *OperationTypeService
	TransactionRepository *repository.TransactionRepository
}

func NewTransactionService() *TransactionService {
	return &TransactionService{
		AccountService:        NewAccountService(),
		OperationTypeService:  NewOperationTypeService(),
		TransactionRepository: repository.NewTransactionRepository(),
	}
}

func (ts *TransactionService) Create(t model.ApiTransaction) (model.Transaction, error) {
	accountDb, err := ts.AccountService.GetById(t.AccountId)
	if err != nil {
		return model.Transaction{}, errors.New("Conta não encontrada")
	}

	operationType, err := ts.OperationTypeService.GetById(t.OperationTypeId)
	if err != nil {
		return model.Transaction{}, errors.New("Operação invalida")
	}

	var newTransaction = model.Transaction{
		Account:       accountDb,
		OperationType: operationType,
		Amount:        t.Amount,
		EventDate:     time.Now(),
	}

	transactionError := newTransaction.CheckTransaction()
	if transactionError != nil {
		return model.Transaction{}, transactionError
	}

	transactionDb, err := ts.TransactionRepository.Create(newTransaction)
	if err != nil {
		return model.Transaction{}, err
	}
	return transactionDb, err
}
