package getGrouptransactions

import domainTransaction "finances.jordis.golang/domain/moves/transactions"

func GetGroupTransactionsQueryHandler(query GetGroupTransactionsQuery, transactionsRepository domainTransaction.TransactionsRepository) ([]*domainTransaction.Transaction, int64, error) {

	transactions, err := transactionsRepository.GetGroupTransactions(query.GroupID, query.FromDate, query.ToDate, query.Category)
	if err != nil {
		return nil, 0, err
	}

	var totalAmount int64
	for _, transaction := range transactions {
		totalAmount += transaction.Amount.Val
	}

	return transactions, totalAmount, err

}
