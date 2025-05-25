package user_transactions

import (
	time_utils "finances.jordis.golang/utils/time"
)

type GetUserTransactionsQuery struct {
	UserID   string
	FromDate time_utils.DateOnly
	ToDate   time_utils.DateOnly
	Category string
}
