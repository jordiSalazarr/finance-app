package userVals

import (
	"errors"
	"strings"
)

const maxLength = 50
const minLength = 2

var (
	ErrInvalidUsernameLength = errors.New("username must be between 5-50 ch")
)

type Username struct {
	Val string
}

func NewUsername(username string) (Username, error) {
	trimmed := strings.Trim(username, " ")
	if len(trimmed) < minLength || len(trimmed) > maxLength {
		return Username{}, ErrInvalidUsernameLength
	}

	return Username{
		Val: trimmed,
	}, nil

}
