package fileprocessor

import (
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
)

type SummaryBuilderSuite struct {
	suite.Suite
	mockCtrl *gomock.Controller
}

func TestSummaryBuilderSuite(t *testing.T) {
	suite.Run(t, new(SummaryBuilderSuite))
}

func (suite *SummaryBuilderSuite) SetupTest() {
	suite.mockCtrl = gomock.NewController(suite.T())
}

func (suite *SummaryBuilderSuite) TearDownTest() {
}

func (suite *SummaryBuilderSuite) TestBuild() {
	suite.Run("given_transactions_when_build_summary_then_return_the_summary_and_no_error", func() {
		builder := newSummaryBuilder()
		defaultDate := time.Date(2022, 12, 18, 12, 0, 0, 0, time.UTC)
		debitA := &transaction{
			date:   defaultDate,
			amount: -60.5,
		}
		debitB := &transaction{
			date:   defaultDate,
			amount: -10.2,
		}
		creditA := &transaction{
			date:   defaultDate,
			amount: 600.65,
		}
		creditB := &transaction{
			date:   defaultDate,
			amount: 44,
		}

		builder.addTransaction(debitA)
		builder.addTransaction(debitB)
		builder.addTransaction(creditA)
		builder.addTransaction(creditB)
		summary := builder.build()

		expectedBalance := debitA.amount + debitB.amount + creditA.amount + creditB.amount
		expectedAverageCredit := (creditA.amount + creditB.amount) / 2
		expectedAverageDebit := (debitA.amount + debitB.amount) / 2
		suite.Require().Equal(expectedBalance, summary.Balance())
		suite.Require().Equal(expectedAverageCredit, summary.AverageCreditAmount())
		suite.Require().Equal(expectedAverageDebit, summary.AverageDebitAmount())
	})

	suite.Run("given_no_transactions_when_build_summary_then_return_the_summary_and_no_error", func() {
		builder := newSummaryBuilder()
		summary := builder.build()

		suite.Require().Equal(0.0, summary.Balance())
		suite.Require().Equal(0.0, summary.AverageCreditAmount())
		suite.Require().Equal(0.0, summary.AverageDebitAmount())
	})
}
