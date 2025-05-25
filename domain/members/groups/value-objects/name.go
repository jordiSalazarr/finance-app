package groupsVals

import (
	"errors"
	"strings"
)

const MinGroupNameLenght = 3
const MaxGroupNameLength = 50

var (
	ErrInvalidGroupNameLength = errors.New("invalid group name length, must be between 3-50 ch")
)

type Name struct {
	Val string
}

func NewGroupName(name string) (Name, error) {
	trimmed := strings.Trim(name, " ")
	if len(trimmed) < MinGroupNameLenght || len(trimmed) > MaxGroupNameLength {
		return Name{}, ErrInvalidGroupNameLength
	}

	return Name{Val: trimmed}, nil

}
