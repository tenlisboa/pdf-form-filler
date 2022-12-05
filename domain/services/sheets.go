package services

import (
	"context"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/tenlisboa/pdf/domain/services/abstract"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
)

type SheetsService struct {
	private       abstract.GoogleAbstractService
	SpreadSheetId string
	Range         string
	Scopes        []string
}

func NewSheetsService(scopes []string, sheetId, rangeName string) SheetsService {
	return SheetsService{
		SpreadSheetId: sheetId,
		Range:         rangeName,
		Scopes:        scopes,
	}
}

func (s SheetsService) GetSheetData() [][]interface{} {
	ctx := context.Background()
	b, err := os.ReadFile("./config/credentials.json")
	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
	}

	// If modifying these scopes, delete your previously saved token.json.
	config, err := google.ConfigFromJSON(b, s.Scopes...)
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
	}

	client := s.private.GetClient(config)

	srv, err := sheets.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		log.Fatalf("Unable to retrieve Sheets client: %v", err)
	}

	response, err := srv.Spreadsheets.Values.Get(s.SpreadSheetId, s.Range).Do()
	if err != nil {
		log.Fatalf("Unable to retrieve data from sheet: %v", err)
	}

	if len(response.Values) == 0 {
		fmt.Println("No data found")
		return nil
	}

	return response.Values
}

func (s SheetsService) GetLastRange() (string, error) {
	f, err := os.ReadFile("./config/last_range")
	if err != nil {
		return "", err
	}

	return string(f), nil
}

func (s SheetsService) SaveLastRange(lastRange string, listLength int) error {
	rg := regexp.MustCompile(`([0-9]+)`)
	numberStr := rg.FindString(lastRange)

	if len(numberStr) < 1 {
		log.Fatalln("Error on saving last range")
	}

	number, err := strconv.ParseInt(numberStr, 0, 64)
	if err != nil {
		return err
	}

	newLastRange := strings.ReplaceAll(lastRange, numberStr, fmt.Sprint(number+int64(listLength)))

	f, err := os.Open("./config/last_range")
	if err != nil {
		f, err = os.Create("./config/last_range")
		if err != nil {
			return err
		}
	}
	defer f.Close()

	_, err = f.WriteString(newLastRange)
	if err != nil {
		return err
	}

	return nil
}
