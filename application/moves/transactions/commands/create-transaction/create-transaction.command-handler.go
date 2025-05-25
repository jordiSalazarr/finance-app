package createtransaction

import (
	"fmt"

	"finances.jordis.golang/domain/members"
	domainTransaction "finances.jordis.golang/domain/moves/transactions"
)

func CreateTransactionCommandHandler(command CreateTransactionsCommand, usersGroupsRepository members.UsersGroupRepository, transactionsRepository domainTransaction.TransactionsRepository) ([]domainTransaction.Transaction, error) {
	transaction, err := domainTransaction.New(command.Description, command.PayedBY,
		command.PayedBY == command.UserID, command.Category, command.Amount, command.Type,
		command.GroupID, command.UserID)

	if err != nil {
		return []domainTransaction.Transaction{}, err
	}

	if ok := isIndividualTransaction(command); ok {
		return handleSingleTransaction(transaction, transactionsRepository)
	}

	transactions := []domainTransaction.Transaction{}

	usersInGroup, err := usersGroupsRepository.GetUsersFromGroup(transaction.GroupId.Val)
	if err != nil {
		return transactions, err
	}
	for _, userID := range usersInGroup {
		var toPayByUser int64 = command.Amount / int64(len(usersInGroup))
		hasAlreadyPayed := userID == command.PayedBY
		if userID == command.PayedBY {
			toPayByUser = command.Amount
		}
		newTransaction, err := domainTransaction.New(command.Description, command.PayedBY,
			hasAlreadyPayed, command.Category, toPayByUser, command.Type,
			command.GroupID, userID)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		transactions = append(transactions, newTransaction)
	}

	err = transactionsRepository.SaveMany(transactions)
	return transactions, err

}

func isIndividualTransaction(command CreateTransactionsCommand) bool {
	return command.GroupID == ""
}

func handleSingleTransaction(transaction domainTransaction.Transaction, transactionsRepository domainTransaction.TransactionsRepository) ([]domainTransaction.Transaction, error) {
	transaction.AlreadyPayed = true
	err := transactionsRepository.SaveOne(transaction)

	return []domainTransaction.Transaction{transaction}, err

}
