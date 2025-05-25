package domainUsers

import "errors"

var (
	ErrNegativeBalanceError = errors.New("negative balance is not allowed")
)
