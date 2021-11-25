package seed

import (
	"pismo-desafio-app/pkg/model"
	"pismo-desafio-app/pkg/repository"
)

type OperationTypeSeeder struct {
	OperationTypeRepository * repository.OperationTypeRepository
}

func NewOperationTypeSeeder() *OperationTypeSeeder  {
	return &OperationTypeSeeder{
		OperationTypeRepository: repository.NewOperationTypeRepository(),
	}
}

func (o *OperationTypeSeeder)RunSeeder() {
	listOperationType, _ := o.OperationTypeRepository.GetAll()

	if len(listOperationType) > 0 {
		return
	}

	var operationTypeDescriptions []string

	operationTypeDescriptions = append(operationTypeDescriptions, "COMPRA A VISTA")
	operationTypeDescriptions = append(operationTypeDescriptions, "COMPRA PARCELADA")
	operationTypeDescriptions = append(operationTypeDescriptions, "SAQUE")
	operationTypeDescriptions = append(operationTypeDescriptions, "PAGAMENTO")

	for i := 0; i < 4; i++ {
		o.OperationTypeRepository.Create(model.OperationType{
			Description: operationTypeDescriptions[i],
		})
	}
}

