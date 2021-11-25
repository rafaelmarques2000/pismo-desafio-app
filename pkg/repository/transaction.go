package repository

import (
	"gorm.io/gorm"
	"pismo-desafio-app/config/database"
	"pismo-desafio-app/pkg/model"
)

type TransactionRepository struct {
	Db *gorm.DB
}

func NewTransactionRepository() *TransactionRepository {
	connection, err := database.OpenConnection()
	if err != nil {
		panic("Falha ao conectar com banco de dados")
	}
	return &TransactionRepository{Db: connection}
}

func (tr *TransactionRepository) Create(t model.Transaction) (model.Transaction, error) {
	result := tr.Db.Create(&t)
	return t, result.Error
}
