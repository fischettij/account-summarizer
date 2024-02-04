package summarizer

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"time"

	"go.uber.org/zap"
)

var (
	ErrUnexpectedHeaders = errors.New("unexpected headers")

	headersFormat = []string{"Id", "Date", "Transaction"}
)

type Logger interface {
	Info(msg string, fields ...zap.Field)
	Error(msg string, fields ...zap.Field)
}

type FileParser struct {
	logger Logger
}

func NewFileParser(logger Logger) (*FileParser, error) {
	if logger == nil {
		return nil, errors.New("logger cannot be nil")
	}
	return &FileParser{logger: logger}, nil
}

func (s *FileParser) GenerateReport(path string) (*Summary, error) {
	builder := newSummaryBuilder()

	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("error opening CSV file: %w", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)

	headers, err := reader.Read()
	if err != nil {
		return nil, fmt.Errorf("error reading CSV headers: %w", err)
	}
	err = validateHeaders(headers)
	if err != nil {
		return nil, err
	}

	// TODO if the id column contains the record id a validation can be added. I assume that id is the transactions ID.
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break // End of File
		}
		if err != nil {
			return nil, fmt.Errorf("error reading CSV record: %w", err)
		}

		date, err := parseDate(record[1])
		if err != nil {
			s.logger.Error(fmt.Sprintf("error on record: %s error parsing date: %s", record, err.Error()))
			return nil, err
		}

		amount, err := strconv.ParseFloat(record[2], 64)
		if err != nil {
			s.logger.Error(fmt.Sprintf("error on record: %s error parsing date: %s", record, err.Error()))
			return nil, err
		}

		trx := &transaction{
			date:   *date,
			amount: amount,
		}

		builder.addTransaction(trx)

		if err != nil {
			return nil, err
		}
	}

	return builder.build(), nil
}

func parseDate(dateString string) (*time.Time, error) {
	// the layout format is M/D. Zero padding is not necessary
	layout := "1/2"

	parsedDate, err := time.Parse(layout, dateString)
	if err != nil {
		return nil, fmt.Errorf("error parsing date: %w", err)
	}

	return &parsedDate, nil

}

// validateHeaders Return error if the headers are different of the expected
// Only validate the first columns. If the file has more columns whit extra data is not a problem.
func validateHeaders(headers []string) error {
	if len(headers) < len(headersFormat) {
		return ErrUnexpectedHeaders
	}

	for i := 0; i < len(headersFormat); i++ {
		if headers[i] != headersFormat[i] {
			return ErrUnexpectedHeaders
		}
	}
	return nil
}
