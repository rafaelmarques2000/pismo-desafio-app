package operationtype

type OperationType struct {
	ID int `json:"id"`
	Description string `json:"description"`
}

func (o *OperationType) GetId() int {
	return o.ID
}

func (o *OperationType) SetId(id int) {
	o.ID = id
}

func (o *OperationType) GetDescription() string {
	return o.Description
}

func (o *OperationType) SetDescription(description string) {
	o.Description = description
}