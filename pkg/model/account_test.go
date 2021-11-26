package tests

import (
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"pismo-desafio-app/pkg/model"
	"testing"
)

//Teste a criação da struct ApiAccount para retorno de dados na api
func TestIfCreateApiAccountStruct(t *testing.T)  {
	var accountMock = model.Account{
		Model: gorm.Model{
			ID:1,
		},
		DocNumber: "123456",
	}
	assert.IsType(t, model.ApiAccount{},accountMock.GetApiAccount())
}

