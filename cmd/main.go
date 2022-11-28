package main

import (
	"github.com/tenlisboa/pdf/domain/usecases"
)

func main() {
	usecase := usecases.FillExpensesFormUsecase{}

	usecase.Execute()
}
