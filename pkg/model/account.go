package model

import "gorm.io/gorm"

type Account struct {
	gorm.Model
	DocNumber string `json:"document_number"`
	AvailableCreditLimit float32 `json:"credit_limit"`
}

type ApiAccount struct {
	Id        uint   `json:"id"`
	DocNumber string `json:"document_number"`
	AvailableCreditLimit float32 `json:"credit_limit"`
}

func (a *Account) GetApiAccount() ApiAccount {
	return ApiAccount{
		Id:        a.Model.ID,
		DocNumber: a.DocNumber,
		AvailableCreditLimit: a.AvailableCreditLimit,
	}
}
