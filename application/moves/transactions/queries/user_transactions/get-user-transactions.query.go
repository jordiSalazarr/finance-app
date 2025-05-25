package user_transactions

import (
	"errors"

	domainTransaction "finances.jordis.golang/domain/moves/transactions"
)

var (
	ErrInvalidDateRange = errors.New("invalid date range")
)

func GetUserTransactionsQueryHandler(query GetUserTransactionsQuery, transactionsRepo domainTransaction.TransactionsRepository) ([]*domainTransaction.Transaction, int64, error) {
	if isValid := query.FromDate.Before(query.ToDate.Time); !isValid {
		return nil, 0, ErrInvalidDateRange

	}

	transactions, err := transactionsRepo.GetUserTransactions(query.UserID, query.FromDate.Time, query.ToDate.Time, query.Category)
	if err != nil {
		return nil, 0, err
	}

	var totalAmount int64
	for _, transaction := range transactions {
		totalAmount += transaction.Amount.Val
	}

	return transactions, totalAmount, nil

}
