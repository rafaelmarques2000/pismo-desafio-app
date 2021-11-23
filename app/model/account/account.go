package account

type Account struct {
	ID int `json:"id"`
	DocNumber string `json:"doc_number"`
}

func (a *Account) GetId() int {
	return a.ID
}

func (a *Account) SetId(id int) {
	a.ID = id
}

func (a *Account) GetDocNumber() string {
	return a.DocNumber
}

func (a *Account) SetDocNumber(docNumber string) {
	a.DocNumber = docNumber
}
