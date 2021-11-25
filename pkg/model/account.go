package model

import "gorm.io/gorm"

type Account struct {
	gorm.Model
	DocNumber string `json:"document_number"`
}

type ApiAccount struct {
	Id        uint   `json:"id"`
	DocNumber string `json:"document_number"`
}

func (a *Account) GetApiAccount() ApiAccount {
	return ApiAccount{
		Id:        a.Model.ID,
		DocNumber: a.DocNumber,
	}
}
