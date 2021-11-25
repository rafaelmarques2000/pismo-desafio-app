package tests

import (
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"pismo-desafio-app/pkg/model"
	"testing"
)

func TestIfCreateApiAccountStruct(t *testing.T)  {
	var accountMock model.Account = model.Account{
		Model: gorm.Model{
			ID:1,
		},
		DocNumber: "123456",
	}
	assert.IsType(t, model.ApiAccount{},accountMock.GetApiAccount())
}

