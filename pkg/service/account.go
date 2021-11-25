package service

import (
	"errors"
	"pismo-desafio-app/pkg/model"
	"pismo-desafio-app/pkg/repository"
)

type AccountService struct {
	AccountRepository *repository.AccountRepository
}

func NewAccountService() *AccountService {
	return &AccountService{
		AccountRepository: repository.NewAccountRepository(),
	}
}

func (as *AccountService)Create(a model.Account) (model.Account, error) {
	listAccounts, _ := as.AccountRepository.GetByDocumentNumber(a.DocNumber)
	if len(listAccounts) > 0 {
		return model.Account{}, errors.New("conta já cadastrada")
	}
	return as.AccountRepository.Create(a)
}

func (as *AccountService) GetById(id uint) (model.Account, error) {
	return as.AccountRepository.GetById(id)
}
