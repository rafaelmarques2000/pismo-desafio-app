package account

import "gorm.io/gorm"

type Account struct {
	gorm.Model
	DocNumber string `json:"doc_number"`
}
