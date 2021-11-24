package operation_type_seeder

import (
	"pismo-desafio-app/app/model/operationtype"
	"pismo-desafio-app/app/repository/operation_type_repository"
)

func RunSeeder() {
	listOperationType, _ := operation_type_repository.GetAll()

	if len(listOperationType) > 0 {
		return
	}

	var operationTypeDescriptions []string

	operationTypeDescriptions = append(operationTypeDescriptions,"COMPRA A VISTA")
	operationTypeDescriptions = append(operationTypeDescriptions,"COMPRA PARCELADA")
	operationTypeDescriptions = append(operationTypeDescriptions,"SAQUE")
	operationTypeDescriptions = append(operationTypeDescriptions,"PAGAMENTO")

	for i := 0; i < 4; i++ {
		operation_type_repository.Create(operationtype.OperationType{
			Description: operationTypeDescriptions[i],
		})
	}
}
