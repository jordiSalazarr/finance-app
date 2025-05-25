package resolve_transaction

import (
	"testing"

	domainTransaction "finances.jordis.golang/domain/moves/transactions"
	inMemoryUsers "finances.jordis.golang/infrastructure/in-memory/members/users"
	trsansactionsInmemory "finances.jordis.golang/infrastructure/in-memory/moves/trsansactions"
)

func TestResolveTransactionCommandHandler(t *testing.T) {

	command := ResolveTransactionCommand{
		TransactionID: "some-id",
		UserID:        "user1",
	}

	transactionsRepo := &trsansactionsInmemory.TransactionsInmemoryRepository{}
	transaction, err := domainTransaction.New("desc", "user2", false, "Alimentación", 200, "SPENDING", "group1", "user1")
	if err != nil {
		t.Error(err.Error())
		return
	}
	mockUsersRepo := &inMemoryUsers.InMemoryUsersRepo{}

	transactionsRepo.SaveOne(transaction)
	err = ResolveTransactionCommandHandler(command, transactionsRepo, mockUsersRepo)
	if err != nil {
		t.Error(err.Error())
		return
	}

	transaction, err = transactionsRepo.GetById(transaction.Pk.Val)
	if err != nil {
		t.Error(err.Error())
		return
	}

	if !transaction.AlreadyPayed {
		t.Error("transaction has not been marked as payed")
	}

}
