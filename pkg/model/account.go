package model

import "gorm.io/gorm"

type Account struct {
	gorm.Model
	DocNumber string `json:"doc_number"`
}

type ApiAccount struct {
	Id        uint   `json:"id"`
	DocNumber string `json:"doc_number"`
}

func (a *Account) GetApiAccount() ApiAccount {
	return ApiAccount{
		Id:        a.Model.ID,
		DocNumber: a.DocNumber,
	}
}
