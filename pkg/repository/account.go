package repository

import (
	"gorm.io/gorm"
	"pismo-desafio-app/config/database"
	"pismo-desafio-app/pkg/model"
)

type AccountRepository struct {
	Db *gorm.DB
}

func NewAccountRepository() *AccountRepository {
	connection, err := database.OpenConnection()
	if err != nil {
		panic("Falha ao conectar com banco de dados")
	}
	return &AccountRepository{Db: connection}
}

func (ar *AccountRepository) Create(a model.Account) (model.Account, error) {
	result := ar.Db.Create(&a)
	return a, result.Error
}

func (ar *AccountRepository) GetById(id uint) (model.Account, error) {
	var accountSearch model.Account
	result := ar.Db.Model(&model.Account{}).First(&accountSearch, id)
	return accountSearch, result.Error
}

func (ar *AccountRepository) GetByDocumentNumber(document string) ([]model.Account, error) {
	var accountSearch []model.Account
	result := ar.Db.Model(&model.Account{}).Where("doc_number = ?", document).Find(&accountSearch)
	return accountSearch, result.Error
}
