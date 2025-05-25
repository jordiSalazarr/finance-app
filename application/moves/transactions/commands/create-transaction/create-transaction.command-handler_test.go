package createtransaction

import (
	"testing"

	usersgorupsInMemory "finances.jordis.golang/infrastructure/in-memory/members/users_gorups"
	trsansactionsInmemory "finances.jordis.golang/infrastructure/in-memory/moves/trsansactions"
	"github.com/stretchr/testify/assert"
)

func TestCreateTransactionCommandHandler(t *testing.T) {
	// Arrange
	mockUsersGroupRepo := &usersgorupsInMemory.UsersgorupsInMemory{}
	mockUsersGroupRepo.UsersGroups = map[string][]string{
		"group1": {"user1", "user2", "user3"},
	}
	inMemoryRepo := &trsansactionsInmemory.TransactionsInmemoryRepository{}

	command := CreateTransactionsCommand{
		Description: "Test Description",
		PayedBY:     "user1",
		UserID:      "user1",
		Category:    "Educación",
		Amount:      100,
		Type:        "SPENDING",
		GroupID:     "group1",
	}
	// Act
	transactions, err := CreateTransactionCommandHandler(command, mockUsersGroupRepo, inMemoryRepo)

	// Assert
	assert.NoError(t, err)
	assert.Len(t, transactions, 3)
	assert.Len(t, inMemoryRepo.Transactions, 3)

	// Verify transaction details
	assert.Equal(t, "Test Transaction", inMemoryRepo.Transactions[0].Description.Val)
	assert.Equal(t, 3333, inMemoryRepo.Transactions[1].Amount.Val)
	assert.Equal(t, "user2", inMemoryRepo.Transactions[1].UserId.Val)
}
