package transaction

import (
	"errors"
	"gorm.io/gorm"
	"pismo-desafio-app/app/model/account"
	"pismo-desafio-app/app/model/operationtype"
	"time"
)

type Transaction struct {
	gorm.Model
	AccountId       int
	Account         account.Account `gorm:"foreignKey:AccountId;references:ID"`
	OperationTypeId int
	OperationType   operationtype.OperationType `gorm:"foreignKey:OperationTypeId;references:ID"`
	Amount          float32
	EventDate       time.Time
}

func (t *Transaction) CheckTransaction() error {
	if t.Amount >= 0 && (
		t.OperationType.ID == operationtype.COMPRA_A_VISTA ||
			t.OperationType.ID == operationtype.COMPRA_PARCELADA ||
			t.OperationType.ID == operationtype.SAQUE) {
		return errors.New("transação invalida para a operação selecionada")
	}
	return nil
}

type ApiTransaction struct {
	Id uint `json:"id"`
	AccountId int `json:"account_id"`
	OperationTypeId int `json:"operation_type_id"`
	Amount float32 `json:"amount"`
	EventDate time.Time `json:"event_date"`
}

