package transaction

import (
	"gorm.io/gorm"
	"pismo-desafio-app/app/model/account"
	"pismo-desafio-app/app/model/operationtype"
	"testing"
)

func TestIfCheckTransactionInvalid(t *testing.T) {

	var tx Transaction = Transaction{
		Amount: 120,
		Account: account.Account{
			Model: gorm.Model{
				ID:1,
			},
		},
		OperationType: operationtype.OperationType{
			Model: gorm.Model{
				ID:operationtype.COMPRA_A_VISTA,
			},
		},
	}

	err := tx.CheckTransaction()

	if err == nil {
		t.Errorf("Transação invalida para teste ")
	}

}

func TestIfCheckTransactionValid(t *testing.T) {

	var tx Transaction = Transaction{
		Amount: -120,
		Account: account.Account{
			Model: gorm.Model{
				ID:1,
			},
		},
		OperationType: operationtype.OperationType{
			Model: gorm.Model{
				ID:operationtype.COMPRA_A_VISTA,
			},
		},
	}

	err := tx.CheckTransaction()

	if err != nil {
		t.Errorf("Transação invalida para teste ")
	}

}