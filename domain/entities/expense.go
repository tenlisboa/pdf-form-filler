package entities

type Expense struct {
	RequestingName     string
	BeneficiaryName    string
	Value              float64
	Description        string
	NFLinks            []string
	AccountBankingData string
	Email              string
	RefundReceived     bool
	FormFilled         bool
	CreatedAt          string
}

func (e Expense) toMap() map[string]interface{} {
	// TODO Terminar, tem alguns campos que vao precisar serem tratados
	mapExpense := map[string]interface{}{
		"solicitante":         e.RequestingName,
		"recebedor":           e.AccountBankingData,
		"descricao_despesa_1": e.Description,
		"valor_total":         e.Value,
	}

	return mapExpense
}
