package transactionVals

import (
	"errors"
	"strings"
)

const maxDescriptionLenght = 50
const minDescriptionLength = 0

var (
	ErrInvalidTransactionDescriptionLength = errors.New("transaction description must be between 0-50 ch")
)

type Description struct {
	Val string
}

func NewDescription(description string) (Description, error) {
	trimmed := strings.Trim(description, " ")
	if len(trimmed) < minDescriptionLength || len(trimmed) > maxDescriptionLenght {
		return Description{}, ErrInvalidTransactionameLength
	}

	return Description{
		Val: trimmed,
	}, nil

}
