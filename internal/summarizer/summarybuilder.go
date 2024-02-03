package summarizer

import "time"

type summaryBuilder struct {
	transactionsByMonth []int

	creditsAmount      float64
	creditTransactions int

	debitsAmount      float64
	debitTransactions int

	totalAmount       float64
	totalTransactions float64
}

func newSummaryBuilder() summaryBuilder {
	return summaryBuilder{
		transactionsByMonth: make([]int, 12),
	}

}

func (b *summaryBuilder) addTransaction(transaction *transaction) {
	b.transactionsByMonth[int(transaction.date.Month())-1]++

	if transaction.amount < 0 {
		b.debitsAmount += transaction.amount
		b.debitTransactions++
	} else {
		b.creditsAmount += transaction.amount
		b.creditTransactions++
	}
	b.totalAmount += transaction.amount
	b.totalTransactions++
}

func (b *summaryBuilder) build() *Summary {
	averageCredits := 0.0
	averageDebits := 0.0
	if b.creditTransactions > 0 {
		averageCredits = b.creditsAmount / float64(b.creditTransactions)
	}
	if b.debitTransactions > 0 {
		averageDebits = b.debitsAmount / float64(b.debitTransactions)
	}

	mapTransactionsByMonth := make(map[time.Month]int)
	for i, transactionCount := range b.transactionsByMonth {
		if transactionCount > 0 {
			mapTransactionsByMonth[time.Month(i+1)] = transactionCount
		}
	}

	return &Summary{
		TransactionsByMonth: mapTransactionsByMonth,
		AverageCreditAmount: averageCredits,
		AverageDebitAmount:  averageDebits,
		Balance:             b.totalAmount,
	}
}
