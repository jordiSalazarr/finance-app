package userVals

import (
	"errors"
	"strings"
)

const minPasswordLength = 7
const maxPasswordLength = 50

var (
	ErrInvalidPassword = errors.New("invalid password, please try again, must be between 7-50 ch")
)

type Password struct {
	Plain  string
	Hashed string
}

type hashService interface {
	Hash(plain string) (string, error)
	Equal(plain string, hashed string) bool
}

// When using this function, a bcrypt object needs to be instantiated and passed along
func NewPassword(plain string, hashService hashService) (Password, error) {
	trimmed := strings.Trim(plain, " ")
	if len(trimmed) < minPasswordLength || len(trimmed) > maxPasswordLength {
		return Password{}, ErrInvalidPassword
	}

	hashed, err := hashService.Hash(trimmed)
	if err != nil {
		return Password{}, err
	}

	return Password{
		Plain:  trimmed,
		Hashed: hashed,
	}, nil

}
