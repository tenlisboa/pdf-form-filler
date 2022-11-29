package entities

type Expense struct {
	Email          string
	RequestingName string
	ReceiverName   string
	RefundType     string
	ExpenseType    string
	Organization   string
	Value          string
	Description    string
	NFLinks        []string
	Observations   string
	CreatedAt      string
}
