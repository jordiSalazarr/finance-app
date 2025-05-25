package userVals

import (
	"fmt"
	"math/rand"
	"time"
)

type VerificationCode struct {
	Val string
}

func NewVerificationCode() VerificationCode {
	rand.Seed(time.Now().UnixNano())
	code := fmt.Sprintf("%06d", rand.Intn(1000000))
	return VerificationCode{Val: code}
}

func ExistingVerificationCode(code string) VerificationCode {

	return VerificationCode{Val: code}
}
