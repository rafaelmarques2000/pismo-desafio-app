package tests

import (
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"pismo-desafio-app/pkg/model"
	"testing"
)

//Verificar se quando o valor for positivo e for uma conta de debito não esta valida
func TestIfCheckInvalidTransactionForOperation(t *testing.T) {
	var tx = model.Transaction{
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

//Verificar se quando o valor for negativo e for uma conta de debito passa no test
func TestIfCheckValiadTransactionForOperation(t *testing.T) {

	var tx = model.Transaction{
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

//Verifica se operações de pagamento não estão sendo lançadas com valor negativo

func TestIfCheckValiadTransactionForOperationPagamento(t *testing.T) {

	var tx = model.Transaction{
		Amount: 120,
		Account: model.Account{
			Model: gorm.Model{
				ID: 1,
			},
		},
		OperationType: model.OperationType{
			Model: gorm.Model{
				ID: model.PAGAMENTO,
			},
		},
	}

	err := tx.CheckTransaction()

	assert.Nil(t, err)
}


//Teste a criação da struct ApiTransaction para retorno de dados na api
func TestIfCreateApiTransactionStruct(t *testing.T)  {
	var transaction= model.Transaction{
		Model: gorm.Model{
			ID:1,
		},
	}
	assert.IsType(t, model.ApiTransaction{},transaction.GetApiTransaction())
}

