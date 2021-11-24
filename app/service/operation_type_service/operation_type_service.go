package operation_type_service

import (
	"pismo-desafio-app/app/model/operationtype"
	"pismo-desafio-app/app/repository/operation_type_repository"
)

func GetById(id int) (operationtype.OperationType, error) {
	return operation_type_repository.GetById(id)
}
