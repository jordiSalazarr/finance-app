package transactionspendingtorecieve

import domainTransaction "finances.jordis.golang/domain/moves/transactions"

func GetTransactionsPendingToRecieve(query GetTransactionsPendingToRecieveQuery, transactionsRepo domainTransaction.TransactionsRepository) ([]*domainTransaction.Transaction, int64, error) {
	transactions, err := transactionsRepo.GetTransactionsPendingToRecieve(query.UserID)
	if err != nil {
		return nil, 0, err
	}

	var total int64
	for _, transaction := range transactions {
		total += transaction.Amount.Val
	}

	return transactions, total, nil
}
