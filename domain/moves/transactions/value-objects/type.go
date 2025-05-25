package transactionVals

import "errors"

type transactionType string

const TransactionTypeIncome transactionType = "INCOME"
const TransactionTypeSpending transactionType = "SPENDING"

var (
	ErrInvalidType = errors.New("invalid transaction type, only INCOME or SPENDING allowed")
)

type Type struct {
	Val transactionType
}

func NewType(typeStr string) (Type, error) {
	inputType := Type{
		Val: transactionType(typeStr),
	}

	if ok := inputType.IsValidType(); !ok {
		return Type{}, ErrInvalidType
	}
	return inputType, nil

}

func (t Type) IsValidType() bool {
	return t.Val == TransactionTypeIncome || t.Val == TransactionTypeSpending
}
