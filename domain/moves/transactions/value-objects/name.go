package transactionVals

import (
	"errors"
	"strings"
)

const maxLength = 50
const minLength = 2

var (
	ErrInvalidTransactionameLength = errors.New("transactioname must be between 2-50 ch")
)

type Name struct {
	Val string
}

func NewTransactioName(transactioname string) (Name, error) {
	trimmed := strings.Trim(transactioname, " ")
	if len(trimmed) < minLength || len(trimmed) > maxLength {
		return Name{}, ErrInvalidTransactionameLength
	}

	return Name{
		Val: trimmed,
	}, nil

}
