package tests

import (
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"pismo-desafio-app/pkg/model"
	"testing"
)

func TestIfCheckTransactionInvalid(t *testing.T) {
	var tx model.Transaction = model.Transaction{
		Amount: 120,
		Account: model.Account{
			Model: gorm.Model{
				ID: 1,
			},
		},
		OperationType: model.OperationType{
			Model: gorm.Model{
				ID: model.COMPRA_A_VISTA,
			},
		},
	}

	err := tx.CheckTransaction()
	assert.NotNil(t, err)
}

func TestIfCheckTransactionValid(t *testing.T) {

	var tx model.Transaction = model.Transaction{
		Amount: -120,
		Account: model.Account{
			Model: gorm.Model{
				ID: 1,
			},
		},
		OperationType: model.OperationType{
			Model: gorm.Model{
				ID: model.COMPRA_A_VISTA,
			},
		},
	}

	err := tx.CheckTransaction()

	assert.Nil(t, err)
}

func TestIfCreateApiTransactionStruct(t *testing.T)  {
	var transactionMock model.Transaction = model.Transaction{
		Model: gorm.Model{
			ID:1,
		},
	}
	assert.IsType(t, model.ApiTransaction{},transactionMock.GetApiTransaction())
}

