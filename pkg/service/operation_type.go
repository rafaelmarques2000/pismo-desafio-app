package service

import (
	"pismo-desafio-app/pkg/model"
	"pismo-desafio-app/pkg/repository"
)

type OperationTypeService struct {
	OperationTypeRepository *repository.OperationTypeRepository
}

func NewOperationTypeService() *OperationTypeService {
	return &OperationTypeService{
		OperationTypeRepository: repository.NewOperationTypeRepository(),
	}
}

func (o *OperationTypeService)GetById(id uint) (model.OperationType, error) {
	return o.OperationTypeRepository.GetById(id)
}