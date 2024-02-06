package summarizer

import (
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/golang/mock/gomock"

	"github.com/fischettij/account-summarizer/pkg/summarizer/mocks"
)

//go:generate mockgen -destination=mocks/mocks.go -package=mocks github.com/fischettij/account-summarizer/pkg/summarizer Summary

func TestGenerateHTMLBody2(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	summary := mocks.NewMockSummary(mockCtrl)

	balance := 99.21
	averageDebit := -71.45
	averageCredit := 128.15
	transactionsByMonth := map[time.Month]int{
		time.January:  10,
		time.February: 5,
	}
	summary.EXPECT().Balance().Return(balance)
	summary.EXPECT().AverageDebitAmount().Return(averageDebit)
	summary.EXPECT().AverageCreditAmount().Return(averageCredit)
	summary.EXPECT().TransactionsByMonth().Return(transactionsByMonth).AnyTimes()

	result, _ := generateHTMLBody(summary)

	expectedSubstrings := []string{
		fmt.Sprintf("<td>Total balance is %.2f</td>", balance),
		fmt.Sprintf("<td>Average debit amount: %.2f</td>", averageDebit),
		fmt.Sprintf("<td>Average credit amount: %.2f</td>", averageCredit),
		fmt.Sprintf("<td>Number of transactions in January: %d</td>", transactionsByMonth[time.January]),
		fmt.Sprintf("<td>Number of transactions in February: %d</td>", transactionsByMonth[time.February]),
	}

	for _, substring := range expectedSubstrings {
		if !strings.Contains(result, substring) {
			t.Errorf("Expected substring %q not found in result", substring)
		}
	}
}
