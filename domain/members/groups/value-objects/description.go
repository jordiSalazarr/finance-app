package groupsVals

import (
	"errors"
	"strings"
)

const MinDescriptionLenght = 1
const MaxDescriptionLength = 500

var (
	ErrInvalidGroupDescriptionLength = errors.New("invalid group description length, must be between 1-500 ch")
)

type Description struct {
	Val string
}

func NewGroupDescription(description string) (Description, error) {
	trimmed := strings.Trim(description, " ")
	if len(trimmed) < MinDescriptionLenght || len(trimmed) > MaxDescriptionLength {
		return Description{}, ErrInvalidGroupDescriptionLength
	}

	return Description{Val: trimmed}, nil

}
