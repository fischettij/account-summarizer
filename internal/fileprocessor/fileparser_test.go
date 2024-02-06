package fileprocessor_test

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"

	"github.com/fischettij/account-summarizer/internal/fileprocessor"
)

type SummarizerSuite struct {
	suite.Suite
	mockCtrl *gomock.Controller
}

func TestMemorySuite(t *testing.T) {
	suite.Run(t, new(SummarizerSuite))
}

func (suite *SummarizerSuite) SetupTest() {
	suite.mockCtrl = gomock.NewController(suite.T())
}

func (suite *SummarizerSuite) TearDownTest() {
}

func (suite *SummarizerSuite) TestGenerateReport() {
	suite.Run("given_a_file_path_with_expected_format_when_parse_it_return_a_summary_and_no_error", func() {
		fileName := "test_file.csv"
		tempFile, err := os.Create(fileName)
		suite.Require().NoError(err)
		defer os.Remove(tempFile.Name())

		content := []string{
			"Id,Date,Transaction",
			"1,7/15,+60.5",
			"1,8/15,+10.0",
		}

		for _, line := range content {
			_, err = fmt.Fprintln(tempFile, line)
			suite.Require().NoError(err)
		}
		tempFile.Close()

		fileParser, err := fileprocessor.NewFileParser(zap.NewNop(), "./")
		suite.Require().NoError(err)

		summary, err := fileParser.GenerateReport(fileName)
		suite.Require().NoError(err)
		suite.Require().NotNil(summary)
		suite.Require().Equal(35.25, summary.AverageCreditAmount())
		suite.Require().Equal(0.0, summary.AverageDebitAmount())
		suite.Require().Equal(70.5, summary.Balance())
		suite.Require().Equal(2, len(summary.TransactionsByMonth()))
		suite.Require().Equal(1, summary.TransactionsByMonth()[time.July])
		suite.Require().Equal(1, summary.TransactionsByMonth()[time.August])
	})

	suite.Run("given_a_file_without_transactions_when_generate_report_then_return_summary_and_no_error", func() {
		fileName := "test_file.csv"
		tempFile, err := os.Create(fileName)
		suite.Require().NoError(err)
		defer os.Remove(tempFile.Name())

		header := "Id,Date,Transaction"
		testData := []byte(fmt.Sprintf("%s\n", header))
		_, err = tempFile.Write(testData)
		suite.Require().NoError(err)
		tempFile.Close()

		fileParser, err := fileprocessor.NewFileParser(zap.NewNop(), "./")
		suite.Require().NoError(err)

		summary, err := fileParser.GenerateReport(fileName)
		suite.Require().NoError(err)
		suite.Require().NotNil(summary)
		suite.Require().Equal(0.0, summary.AverageCreditAmount())
		suite.Require().Equal(0.0, summary.AverageDebitAmount())
		suite.Require().Equal(0.0, summary.Balance())
	})

	suite.Run("given_a_file_wit_invalid_headers_when_generate_report_then_return_error", func() {
		fileName := "test_file.csv"
		tempFile, err := os.Create(fileName)
		suite.Require().NoError(err)
		defer os.Remove(tempFile.Name())

		invalidHeader := "invalid,header,error"
		testData := []byte(fmt.Sprintf("%s\n", invalidHeader))
		_, err = tempFile.Write(testData)
		suite.Require().NoError(err)
		tempFile.Close()

		fileParser, err := fileprocessor.NewFileParser(zap.NewNop(), "./")
		suite.Require().NoError(err)

		summary, err := fileParser.GenerateReport(fileName)
		suite.Require().Error(err)
		suite.Require().ErrorIs(err, fileprocessor.ErrUnexpectedHeaders)
		suite.Require().Nil(summary)
	})

}
