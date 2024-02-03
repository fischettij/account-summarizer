package summarizer

import "time"

type transaction struct {
	date   time.Time
	amount float64
}

type Summary struct {
	TransactionsByMonth map[time.Month]int
	AverageCreditAmount float64
	AverageDebitAmount  float64
	Balance             float64
}
