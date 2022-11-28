package usecases

import (
	"fmt"
	"strconv"
	"strings"

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
		"16heR7XYYqV_1Bq8087syWbd3q0_NrkSbQUmqFgcnOGE",
		"A2:J",
	)

	expenses := presentData(sv.GetSheetData())

	fmt.Println(expenses)
}

func presentData(values [][]interface{}) []entities.Expense {
	expenses := []entities.Expense{}
	for _, row := range values {
		parsedValue, _ := strconv.ParseFloat(fmt.Sprintf("%s", row[3]), 64)
		expenses = append(expenses, entities.Expense{
			RequestingName:     fmt.Sprintf("%s", row[1]),
			BeneficiaryName:    fmt.Sprintf("%s", row[2]),
			Value:              parsedValue,
			Description:        fmt.Sprintf("%s", row[4]),
			NFLinks:            strings.Split(fmt.Sprintf("%s", row[5]), ", "),
			AccountBankingData: fmt.Sprintf("%s", row[6]),
			Email:              fmt.Sprintf("%s", row[7]),
			RefundReceived:     fmt.Sprintf("%s", row[8]) == "Sim",
			FormFilled:         fmt.Sprintf("%s", row[9]) == "Sim",
			CreatedAt:          fmt.Sprintf("%s", row[0]),
		})
	}

	return expenses
}
