package transactionVals

import "errors"

const maxAmount = 2147483647
const minAmount = -2147483648

var (
	ErrAmount = errors.New("max limit reached, please try with a lower number")
)

type Amount struct {
	Val int64
}

func NewAmount(val int64) (Amount, error) {
	if val >= maxAmount || val <= minAmount {
		return Amount{}, ErrAmount
	}

	return Amount{Val: val}, nil

}
