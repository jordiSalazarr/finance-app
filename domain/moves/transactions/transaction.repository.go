package domainTransaction

import (
	"time"
)

type TransactionsRepository interface {
	SaveMany(transaction []Transaction) error
	SaveOne(transaction Transaction) error
	GetById(id string) (Transaction, error)
	MarkAsPayed(id string) error
	GetUserTransactions(userID string, from time.Time, to time.Time, category string) ([]*Transaction, error)
	GetGroupTransactions(groupID string, fromDate time.Time, toDate time.Time, category string) ([]*Transaction, error)
	GetTransactionsPendingToRecieve(userID string) ([]*Transaction, error)
	GetTransactionsPendingToPay(userID string) ([]*Transaction, error)
}
