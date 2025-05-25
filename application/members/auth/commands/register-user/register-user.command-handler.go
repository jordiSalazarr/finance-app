package auth_useCases

import (
	"errors"

	domainUsers "finances.jordis.golang/domain/members/users"
	hashService "finances.jordis.golang/services"
	mail_service "finances.jordis.golang/services/mail"
)

var (
	ErrInvalidInput = errors.New("invalid user input data, please try again with different parameters")
)

// TODO: Mail service should be an interface
func RegisterUserCommandHandler(command RegisterUserCommand, userRepository domainUsers.UserRepository, mailSerivce *mail_service.SMPTService) (domainUsers.User, error) {
	if exists := userRepository.Exists(command.Mail); exists {
		return domainUsers.User{}, ErrInvalidInput

	}
	bcryptService := &hashService.BCrypt{}
	user, err := domainUsers.NewUser(command.Name, command.MonthlyIncome, command.Mail, command.Password, command.CurrentBalance, bcryptService)
	if err != nil {
		return domainUsers.User{}, err
	}

	err = userRepository.Save(user)
	if err != nil {
		return domainUsers.User{}, err
	}

	go mailSerivce.SendVerificationCode(command.Mail, user.VerificationCode.Val)

	return user, nil

}
