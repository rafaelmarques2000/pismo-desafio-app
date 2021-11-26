package model

import (
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"testing"
)

//Verificar se quando o valor for positivo e for uma conta de debito não esta valida
func TestIfCheckInvalidTransactionForOperation(t *testing.T) {
	var tx = Transaction{
		Amount: 120,
		Account: Account{
			Model: gorm.Model{
				ID: 1,
			},
		},
		OperationType: OperationType{
			Model: gorm.Model{
				ID: COMPRA_A_VISTA,
			},
		},
	}

	err := tx.CheckTransaction()
	assert.NotNil(t, err)
}

//Verificar se quando o valor for negativo e for uma conta de debito passa no test
func TestIfCheckValiadTransactionForOperation(t *testing.T) {

	var tx = Transaction{
		Amount: -120,
		Account: Account{
			Model: gorm.Model{
				ID: 1,
			},
		},
		OperationType: OperationType{
			Model: gorm.Model{
				ID: COMPRA_A_VISTA,
			},
		},
	}

	err := tx.CheckTransaction()

	assert.Nil(t, err)
}

//Verifica se operações de pagamento não estão sendo lançadas com valor negativo

func TestIfCheckValiadTransactionForOperationPagamento(t *testing.T) {

	var tx = Transaction{
		Amount: 120,
		Account: Account{
			Model: gorm.Model{
				ID: 1,
			},
		},
		OperationType: OperationType{
			Model: gorm.Model{
				ID: PAGAMENTO,
			},
		},
	}

	err := tx.CheckTransaction()

	assert.Nil(t, err)
}


//Teste a criação da struct ApiTransaction para retorno de dados na api
func TestIfCreateApiTransactionStruct(t *testing.T)  {
	var transaction= Transaction{
		Model: gorm.Model{
			ID:1,
		},
	}
	assert.IsType(t, ApiTransaction{},transaction.GetApiTransaction())
}

