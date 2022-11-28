package services

import (
	"log"

	"github.com/desertbit/fillpdf"
)

type PDFService struct{}

func (pdf PDFService) FillForm(data map[string]interface{}, formPath, filledFormPath string) {
	// Fill the form PDF with our values.
	err := fillpdf.Fill(data, formPath, filledFormPath)
	if err != nil {
		log.Fatal(err)
	}
}
