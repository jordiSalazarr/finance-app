package auth_useCases

import (
	"errors"

	domainUsers "finances.jordis.golang/domain/members/users"
)

var (
	ErrInvalidData = errors.New("invalid data input")
)

func VerifyUserCommandHandler(command VerifyUserCommand, userRepository domainUsers.UserRepository) error {

	user, err := userRepository.GetUser(command.Email)
	if err != nil {
		return err
	}
	if user.VerificationCode.Val != command.Code {
		return ErrInvalidData
	}

	return userRepository.VerificateUser(command.Email)

}
