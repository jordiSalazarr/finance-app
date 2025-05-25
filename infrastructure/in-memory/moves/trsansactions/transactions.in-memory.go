package trsansactionsInmemory

import (
	"errors"
	"fmt"
	"time"

	domainTransaction "finances.jordis.golang/domain/moves/transactions"
)

type TransactionsInmemoryRepository struct {
	Transactions []*domainTransaction.Transaction
}

func New() *TransactionsInmemoryRepository {
	return &TransactionsInmemoryRepository{
		Transactions: []*domainTransaction.Transaction{},
	}
}
func (im *TransactionsInmemoryRepository) SaveOne(transaction domainTransaction.Transaction) error {
	im.Transactions = append(im.Transactions, &transaction)
	return nil
}

func (im *TransactionsInmemoryRepository) SaveMany(transactions []domainTransaction.Transaction) error {
	for _, transaction := range transactions {
		im.Transactions = append(im.Transactions, &transaction)

	}

	return nil
}

func (im *TransactionsInmemoryRepository) GetById(id string) (domainTransaction.Transaction, error) {
	for _, transaction := range im.Transactions {
		if transaction.Pk.Val == id {
			return *transaction, nil
		}
	}

	return domainTransaction.Transaction{}, errors.New("transaction not found")
}

func (im *TransactionsInmemoryRepository) MarkAsPayed(id string) error {
	for _, transaction := range im.Transactions {
		if transaction.Pk.Val == id {
			transaction.AlreadyPayed = true
			return nil
		}
	}

	return errors.New("transaction not found")
}

func (im *TransactionsInmemoryRepository) GetUserTransactions(userID string, from time.Time, to time.Time, catgeory string) ([]*domainTransaction.Transaction, error) {
	fmt.Print("transactions in mem: ", im.Transactions)
	var transactions []*domainTransaction.Transaction
	for _, transaction := range im.Transactions {
		if transaction.UserId.Val == userID && transaction.CreatedAt.After(from) && transaction.CreatedAt.Before(to) && transaction.Category.Val == catgeory || catgeory == "" {
			transactions = append(transactions, transaction)
		}
	}

	return transactions, nil

}
func (im *TransactionsInmemoryRepository) GetGroupTransactions(groupID string, fromDate time.Time, toDate time.Time, category string) ([]*domainTransaction.Transaction, error) {
	return im.Transactions, nil

}

func (im *TransactionsInmemoryRepository) GetTransactionsPendingToPay(userID string) ([]*domainTransaction.Transaction, error) {
	var transactions []*domainTransaction.Transaction
	for _, transaction := range im.Transactions {
		if transaction.UserId.Val == userID && !transaction.AlreadyPayed {
			transactions = append(transactions, transaction)
		}
	}

	if len(transactions) == 0 {
		return nil, errors.New("no pending transactions found")
	}

	return transactions, nil
}

func (im *TransactionsInmemoryRepository) GetTransactionsPendingToRecieve(userID string) ([]*domainTransaction.Transaction, error) {
	var transactions []*domainTransaction.Transaction
	for _, transaction := range im.Transactions {
		if transaction.UserId.Val != userID && !transaction.AlreadyPayed {
			transactions = append(transactions, transaction)
		}
	}

	if len(transactions) == 0 {
		return nil, errors.New("no pending transactions found")
	}

	return transactions, nil
}
