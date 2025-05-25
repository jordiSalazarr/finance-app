package hashService

import (
	"golang.org/x/crypto/bcrypt"
)

type BCrypt struct{}

func (b *BCrypt) Hash(plain string) (string, error) {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(plain), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedBytes), nil
}

func (b *BCrypt) Equal(plain string, hashed string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(plain))
	return err == nil
}
