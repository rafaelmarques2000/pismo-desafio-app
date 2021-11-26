package model

import (
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"testing"
)

//Teste a criação da struct ApiAccount para retorno de dados na api
func TestIfCreateApiAccountStruct(t *testing.T)  {
	var accountMock = Account{
		Model: gorm.Model{
			ID:1,
		},
		DocNumber: "123456",
	}
	assert.IsType(t, ApiAccount{},accountMock.GetApiAccount())
}

