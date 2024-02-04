package fileprocessor

import "time"

type transaction struct {
	date   time.Time
	amount float64
}

type Summary struct {
	transactionsByMonth map[time.Month]int
	averageCreditAmount float64
	averageDebitAmount  float64
	balance             float64
}

func (s *Summary) Balance() float64 {
	return s.balance
}

func (s *Summary) AverageCreditAmount() float64 {
	return s.averageCreditAmount
}
func (s *Summary) AverageDebitAmount() float64 {
	return s.averageDebitAmount
}
