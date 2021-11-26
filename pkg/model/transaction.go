package model

import (
	"errors"
	"gorm.io/gorm"
	"time"
)

/** ORM model e regra de negocio de validação para a transação **/
type Transaction struct {
	gorm.Model
	AccountId       int
	Account         Account `gorm:"foreignKey:AccountId;references:ID"`
	OperationTypeId int
	OperationType   OperationType `gorm:"foreignKey:OperationTypeId;references:ID"`
	Amount          float32
	EventDate       time.Time
}

/** Struct Simples para retorno da api e entrada de dados **/
type ApiTransaction struct {
	Id              uint      `json:"id"`
	AccountId       uint      `json:"account_id"`
	OperationTypeId uint      `json:"operation_type_id"`
	Amount          float32   `json:"amount"`
	EventDate       time.Time `json:"event_date"`
}

//Metodo para verificar se a transação é válida
func (t *Transaction) CheckTransaction() error {
	if t.Amount >= 0 && (t.OperationType.ID == COMPRA_A_VISTA ||
		t.OperationType.ID == COMPRA_PARCELADA ||
		t.OperationType.ID == SAQUE) {
		return errors.New("transação invalida Operações de Compra a vista, Compra parcelada onde o valor deve ser debitado")
	}

	if t.Amount < 0 && t.OperationType.ID == PAGAMENTO {
		return errors.New("transação invalida")
	}

	return nil
}

//Metodo para criar uma versão mais simples do modelo para retorno via http
func (t *Transaction) GetApiTransaction() ApiTransaction {
	return ApiTransaction{
		Id:              t.Model.ID,
		AccountId:       t.Account.ID,
		OperationTypeId: t.OperationType.ID,
		Amount:          t.Amount,
		EventDate:       t.EventDate,
	}
}
