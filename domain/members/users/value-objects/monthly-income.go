package userVals

import "errors"

const (
	MAX_MONTHLY_INCOME = 100000000
)

var (
	ErrInvalidMonthlyIncome = errors.New("invalid montlhy income value")
)

type MonthlyIncome struct {
	Val int64
}

func NewMonthlyIncome(val int64) (MonthlyIncome, error) {
	if val <= 0 || val >= MAX_MONTHLY_INCOME {
		return MonthlyIncome{}, ErrInvalidMonthlyIncome
	}

	return MonthlyIncome{Val: val}, nil
}
