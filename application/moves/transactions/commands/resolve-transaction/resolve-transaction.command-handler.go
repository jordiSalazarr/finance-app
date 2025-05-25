package resolve_transaction

import (
	domainUsers "finances.jordis.golang/domain/members/users"
	domainTransaction "finances.jordis.golang/domain/moves/transactions"
)

func ResolveTransactionCommandHandler(command ResolveTransactionCommand, transactionsRepository domainTransaction.TransactionsRepository, usersRepository domainUsers.UserRepository) error {
	transaction, err := transactionsRepository.GetById(command.TransactionID)
	if err != nil {
		return err
	}
	err = transactionsRepository.MarkAsPayed(command.TransactionID)
	if err != nil {
		return err
	}
	return usersRepository.UpdateActorsCurrentBalance(command.UserID, transaction.PayedBy.Val, transaction.Amount.Val)

}
