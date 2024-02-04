package summarizer

import (
	"context"
	"errors"
)

type Summarizer interface {
	GenerateReport(fileName string) (Summary, error)
}

type Summary interface {
	Balance() float64
	AverageCreditAmount() float64
	AverageDebitAmount() float64
}

type Manager struct {
	summarizer Summarizer
}

func NewManager(summarizer Summarizer) (*Manager, error) {
	if summarizer == nil {
		return nil, errors.New("summarizer cannot be nil")
	}

	return &Manager{
		summarizer: summarizer,
	}, nil
}

func (m *Manager) SendSummary(ctx context.Context, _, fileName string) error {
	_, err := m.summarizer.GenerateReport(fileName)
	return err
}
