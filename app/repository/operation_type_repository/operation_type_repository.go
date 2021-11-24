package operation_type_repository

import (
	"pismo-desafio-app/app/model/operationtype"
	"pismo-desafio-app/config/database"
)

func GetAll() ([]operationtype.OperationType, error) {
	connection, err := database.OpenConnection()
	if err != nil {
		return []operationtype.OperationType{}, err
	}

	var listOperationType []operationtype.OperationType
	result := connection.Model(&operationtype.OperationType{}).Find(&listOperationType)
	return listOperationType, result.Error
}

func Create(OperationType operationtype.OperationType) error {
	connection, err := database.OpenConnection()
	if err != nil {
		return err
	}
	result := connection.Create(&OperationType)
	return result.Error
}

func GetById(id int) (operationtype.OperationType, error) {
	connection, err := database.OpenConnection()
	if err != nil {
		return operationtype.OperationType{}, err
	}
	var OperationTypeSearch operationtype.OperationType
	result := connection.Model(&operationtype.OperationType{}).First(&OperationTypeSearch, id)
	return OperationTypeSearch, result.Error
}
