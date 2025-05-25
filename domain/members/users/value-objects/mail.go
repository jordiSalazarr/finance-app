package userVals

import (
	"errors"
	"regexp"
)

var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)
var (
	ErrInvalidMail = errors.New("invalid mail, please try again")
)

type Mail struct {
	Val string
}

func NewMail(mail string) (Mail, error) {
	if ok := emailRegex.MatchString(mail); !ok {
		return Mail{}, ErrInvalidMail
	}

	return Mail{Val: mail}, nil

}
