package api

import (
	"net/http"
	"time"

	createtransaction "finances.jordis.golang/application/moves/transactions/commands/create-transaction"
	resolve_transaction "finances.jordis.golang/application/moves/transactions/commands/resolve-transaction"
	getGrouptransactions "finances.jordis.golang/application/moves/transactions/queries/group_transactions"
	transactionspendingtopay "finances.jordis.golang/application/moves/transactions/queries/pending_to_pay"
	transactionspendingtorecieve "finances.jordis.golang/application/moves/transactions/queries/pending_to_recieve"
	"finances.jordis.golang/application/moves/transactions/queries/user_transactions"
	domainTransaction "finances.jordis.golang/domain/moves/transactions"
	time_utils "finances.jordis.golang/utils/time"
	"github.com/gin-gonic/gin"
)

type TransactionDTO struct {
	Pk           string    `json:"pk"`
	Description  string    `json:"description"`
	Category     string    `json:"category"`
	Amount       int64     `json:"amount"`
	Type         string    `json:"type"`
	GroupId      string    `json:"groupId"`
	UserId       string    `json:"userId"`
	CreatedAt    time.Time `json:"createdAt"`
	PayedBy      string    `json:"payedBy"`
	AlreadyPayed bool      `json:"alreadyPayed"`
}

func (app *App) CreateTransactionHandler(c *gin.Context) {
	userId, exists := GetUserIdFromRequest(c)
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "User ID not found in request",
		})
		return
	}
	var input struct {
		Description string `json:"description"`
		Amount      int64  `json:"amount"`
		Type        string `json:"type"`
		GroupID     string `json:"group_id"`
		Category    string `json:"category"`
	}

	err := c.BindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid input",
		})
		return
	}

	command := createtransaction.CreateTransactionsCommand{
		Description: input.Description,
		Amount:      input.Amount,
		Type:        input.Type,
		GroupID:     input.GroupID,
		Category:    input.Category,
		PayedBY:     userId,
		UserID:      userId,
	}

	_, err = createtransaction.CreateTransactionCommandHandler(command, app.UsersGroupsRepo, app.TransactionsRepo, app.UsersRepo)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"succes": "transaction created succesfully",
	})
}

func (app *App) ResolveTransaction(c *gin.Context) {
	userId, exists := GetUserIdFromRequest(c)
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "User ID not found in request",
		})
		return
	}
	transactionID := c.Params.ByName("transactionID")
	if transactionID == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "bad input, cannot find transaction id",
		})
		return
	}
	command := resolve_transaction.ResolveTransactionCommand{
		TransactionID: transactionID,
		UserID:        userId,
	}
	err := resolve_transaction.ResolveTransactionCommandHandler(command, app.TransactionsRepo, app.UsersRepo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Transaction marked as payed",
	})
}

func (app *App) GetUserTransactions(c *gin.Context) {
	userId, exists := GetUserIdFromRequest(c)
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "User ID not found in request",
		})
		return
	}
	var input struct {
		FromDate time_utils.DateOnly `form:"from_date"`
		ToDate   time_utils.DateOnly `form:"to_date"`
		Category string              `form:"category"`
	}

	err := c.BindQuery(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	if input.Category == "" {
		input.Category = "all"
	}

	query := user_transactions.GetUserTransactionsQuery{
		UserID:   userId,
		FromDate: input.FromDate,
		ToDate:   input.ToDate,
		Category: input.Category,
	}

	transactions, totalAmount, err := user_transactions.GetUserTransactionsQueryHandler(query, app.TransactionsRepo)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	transactionsResponse := mapToTransactionDTO(transactions)

	c.JSON(http.StatusOK, gin.H{
		"transactions": transactionsResponse,
		"total_amount": totalAmount,
	})

}

func (app *App) GetGroupTransactions(c *gin.Context) {
	groupId := c.Param("group_id")
	var input struct {
		FromDate time_utils.DateOnly `form:"from_date"`
		ToDate   time_utils.DateOnly `form:"to_date"`
		Category string              `form:"category"`
	}

	err := c.BindQuery(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": err.Error(),
		})
		return
	}

	query := getGrouptransactions.GetGroupTransactionsQuery{
		GroupID:  groupId,
		FromDate: input.FromDate.In(time.UTC),
		ToDate:   input.ToDate.In(time.UTC),
		Category: input.Category,
	}

	transactions, totalAmount, err := getGrouptransactions.GetGroupTransactionsQueryHandler(query, app.TransactionsRepo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"err": err.Error(),
		})
		return
	}

	response := mapToTransactionDTO(transactions)

	c.JSON(http.StatusOK, gin.H{
		"transactions": response,
		"total_amount": totalAmount,
	})

}

func (app *App) GetPendingToRecieveTransactions(c *gin.Context) {
	userId, exists := GetUserIdFromRequest(c)
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "User ID not found in request",
		})
		return
	}
	query := transactionspendingtorecieve.GetTransactionsPendingToRecieveQuery{
		UserID: userId,
	}

	transactions, totalToRecieve, err := transactionspendingtorecieve.GetTransactionsPendingToRecieve(query, app.TransactionsRepo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"err": err.Error(),
		})
		return
	}

	response := mapToTransactionDTO(transactions)
	c.JSON(http.StatusOK, gin.H{
		"transactions": response,
		"total_amount": totalToRecieve,
	})

}

func (app *App) GetPendingToPayTransactions(c *gin.Context) {
	userId, exists := GetUserIdFromRequest(c)
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "User ID not found in request",
		})
		return
	}
	query := transactionspendingtopay.GetTransactionsPendingToPayQuery{
		UserID: userId,
	}

	transactions, totalToPay, err := transactionspendingtopay.GetTransactionsPendingToPay(query, app.TransactionsRepo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"err": err.Error(),
		})
		return
	}

	response := mapToTransactionDTO(transactions)

	c.JSON(http.StatusOK, gin.H{
		"transactions": response,
		"total_amount": totalToPay,
	})

}

func mapToTransactionDTO(transactions []*domainTransaction.Transaction) []*TransactionDTO {
	var transactionDTOs []*TransactionDTO
	for _, transaction := range transactions {
		transactionDTO := &TransactionDTO{
			Pk:           transaction.Pk.Val,
			Description:  transaction.Description.Val,
			Category:     transaction.Category.Val,
			Amount:       transaction.Amount.Val,
			Type:         string(transaction.Type.Val),
			GroupId:      transaction.GroupId.Val,
			UserId:       transaction.UserId.Val,
			CreatedAt:    transaction.CreatedAt,
			PayedBy:      transaction.PayedBy.Val,
			AlreadyPayed: transaction.AlreadyPayed,
		}
		transactionDTOs = append(transactionDTOs, transactionDTO)
	}

	return transactionDTOs
}
