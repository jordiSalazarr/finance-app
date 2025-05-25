package getGrouptransactions

import (
	"time"
)

type GetGroupTransactionsQuery struct {
	GroupID  string
	FromDate time.Time
	ToDate   time.Time
	Category string
}
