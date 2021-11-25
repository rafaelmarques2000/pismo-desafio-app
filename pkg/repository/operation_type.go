package repository

import (
	"gorm.io/gorm"
	"pismo-desafio-app/config/database"
	"pismo-desafio-app/pkg/model"
)

type OperationTypeRepository struct {
	Db *gorm.DB
}

func NewOperationTypeRepository() *OperationTypeRepository {
	connection, err := database.OpenConnection()
	if err != nil {
		panic("Falha ao conectar com banco de dados")
	}
	return &OperationTypeRepository{Db: connection}
}

func (o *OperationTypeRepository) GetAll() ([]model.OperationType, error) {
	var listOperationType []model.OperationType
	result := o.Db.Model(&model.OperationType{}).Find(&listOperationType)
	return listOperationType, result.Error
}

func (o *OperationTypeRepository) Create(OperationType model.OperationType) error {
	result := o.Db.Create(&OperationType)
	return result.Error
}

func (o *OperationTypeRepository) GetById(id uint) (model.OperationType, error) {
	var OperationTypeSearch model.OperationType
	result := o.Db.Model(&model.OperationType{}).First(&OperationTypeSearch, id)
	return OperationTypeSearch, result.Error
}
