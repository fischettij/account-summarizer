package summarizer

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"
)

const emailBalanceSubject = "Hi! This is your annual balance"

type Summarizer interface {
	GenerateReport(fileName string) (Summary, error)
}

type EmailSender interface {
	SendEmail(to string, message []byte) error
}

type Summary interface {
	Balance() float64
	AverageCreditAmount() float64
	AverageDebitAmount() float64
	TransactionsByMonth() map[time.Month]int
}

type Manager struct {
	summarizer  Summarizer
	emailSender EmailSender
}

func NewManager(summarizer Summarizer, emailSender EmailSender) (*Manager, error) {
	if summarizer == nil {
		return nil, errors.New("summarizer cannot be nil")
	}
	if emailSender == nil {
		return nil, errors.New("emailSender cannot be nil")
	}

	return &Manager{
		summarizer:  summarizer,
		emailSender: emailSender,
	}, nil
}

func (m *Manager) SendSummary(ctx context.Context, emailTo, fileName string) error {
	summary, err := m.summarizer.GenerateReport(fileName)
	if err != nil {
		return err
	}

	message := fmt.Sprintf("Subject: %s\r\n\r\n%s", emailBalanceSubject, generateBody(summary))
	err = m.emailSender.SendEmail(emailTo, []byte(message))
	return err
}

func generateBody(summary Summary) string {
	header := "Hi! This is your anual account summary."
	totalBalance := fmt.Sprintf("Total balance is %.2f", summary.Balance())
	debitAverage := fmt.Sprintf("Average debit amount: %.2f", summary.AverageDebitAmount())
	creditAverage := fmt.Sprintf("Average credit amount: %.2f", summary.AverageCreditAmount())
	body := strings.Join([]string{header, totalBalance, debitAverage, creditAverage}, "\n")

	for month := time.January; month <= time.December; month++ {
		value, found := summary.TransactionsByMonth()[month]
		if found && value > 0 {
			body += fmt.Sprintf("\nNumber of transactions in %s: %d", month.String(), value)
		}
	}

	return body
}
