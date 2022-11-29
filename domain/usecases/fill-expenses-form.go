package usecases

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/tenlisboa/pdf/domain/entities"
	"github.com/tenlisboa/pdf/domain/services"
)

type FillExpensesFormUsecase struct{}

func (uc FillExpensesFormUsecase) Execute() {
	sv := services.NewSheetsService(
		[]string{
			"https://www.googleapis.com/auth/spreadsheets.readonly",
			"https://www.googleapis.com/auth/drive.readonly",
		},
		"1TwEFNyrAtzjUkZyIvHX3otJGpHgBjIcjFCE02CXECG4",
		"A2:K",
	)

	expenses := dataToEntity(sv.GetSheetData())

	fmt.Println(expenses)

	pdfService := services.NewPDFService()

	for _, expense := range expenses {
		pdfService.FillForm(entityToMap(expense), "./pdf/form.pdf", fmt.Sprintf("./pdf/filled/%s", getPdfNameBasedOnEntity(expense)))
	}
}

func dataToEntity(values [][]interface{}) []entities.Expense {
	expenses := []entities.Expense{}
	for _, row := range values {
		expenses = append(expenses, entities.Expense{
			Email:          fmt.Sprintf("%s", row[1]),
			RequestingName: fmt.Sprintf("%s", row[2]),
			ReceiverName:   fmt.Sprintf("%s", row[3]),
			RefundType:     fmt.Sprintf("%s", row[10]),
			ExpenseType:    fmt.Sprintf("%s", row[4]),
			Organization:   fmt.Sprintf("%s", row[5]),
			Value:          fmt.Sprintf("%s", row[6]),
			Description:    fmt.Sprintf("%s", row[7]),
			NFLinks:        strings.Split(fmt.Sprintf("%s", row[8]), ", "),
			Observations:   fmt.Sprintf("%s", row[9]),
			CreatedAt:      fmt.Sprintf("%s", row[0]),
		})
	}

	return expenses
}

func entityToMap(e entities.Expense) map[string]interface{} {

	dataArr := strings.Split(e.CreatedAt, " ")
	dataArr = strings.Split(dataArr[0], "/")
	mapExpense := map[string]interface{}{
		"solicitante":         e.RequestingName,
		"recebedor":           e.ReceiverName,
		"descricao_despesa_1": e.Description,
		"valor_total":         e.Value,
		"data_dia":            dataArr[0],
		"data_mes":            dataArr[1],
		"data_ano":            dataArr[2],
		"p_lider":             getStrBasedOnRefundType(e.RefundType, "p_lider"),
		"p_fornecedor":        getStrBasedOnRefundType(e.RefundType, "p_fornecedor"),
		"p_adiantamento":      getStrBasedOnRefundType(e.RefundType, "p_adiantamento"),
		"orcamento_valor":     getValueBasedOnExpenseType(e.ExpenseType, "orcamento_valor", e.Value),
		"jejum_alim_vest":     getValueBasedOnExpenseType(e.ExpenseType, "jejum_alim_vest", e.Value),
		"jejum_med":           getValueBasedOnExpenseType(e.ExpenseType, "jejum_med", e.Value),
		"jejum_moradia":       getValueBasedOnExpenseType(e.ExpenseType, "jejum_moradia", e.Value),
		"jejum_servico_pub":   getValueBasedOnExpenseType(e.ExpenseType, "jejum_servico_pub", e.Value),
		"conta_cadastrada":    "X",
	}

	return mapExpense
}

func getStrBasedOnRefundType(refundType, field string) string {
	if strings.EqualFold(refundType, field) {
		return "X"
	}

	return ""
}

func getValueBasedOnExpenseType(expenseType, value, field string) string {
	fastField := map[string]string{
		"alimentacao_vestimenta": "jejum_alim_vest",
		"moradia":                "jejum_moradia",
		"despesa_medica":         "jejum_med",
		"servicos_publicos":      "jejum_servico_pub",
		"orcamento":              "orcamento_valor",
	}

	if fastField[expenseType] == field {
		return value
	}

	return ""
}

func getPdfNameBasedOnEntity(expense entities.Expense) string {
	dateTime, err := time.Parse("02/01/2006 15:04:05", expense.CreatedAt)
	if err != nil {
		log.Fatalln(err)
	}
	return fmt.Sprintf("%d_%s.pdf", dateTime.Unix(), strings.ReplaceAll(strings.ToLower(expense.RequestingName), " ", "_"))
}
