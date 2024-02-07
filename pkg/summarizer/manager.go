package summarizer

import (
	"context"
	"errors"
	"fmt"
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
		return nil, errors.New("summarizer can not be nil")
	}
	if emailSender == nil {
		return nil, errors.New("emailSender can not be nil")
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

	mimeType := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	body, err := GenerateHTMLBody(summary)
	if err != nil {
		return err
	}
	// TODO: this logic about the emails message format could be in email package. Its implementation logic about how emails is sent
	message := fmt.Sprintf("Subject: %s\r\n%s%s", emailBalanceSubject, mimeType, body)
	err = m.emailSender.SendEmail(emailTo, []byte(message))
	return err
}
