package auth_useCases

import (
	"errors"

	domainUsers "finances.jordis.golang/domain/members/users"
	hashService "finances.jordis.golang/services"
	jwtService "finances.jordis.golang/services/jwt"
)

var (
	ErrInvalidData = errors.New("invalid data input, please try again")
)

func LoginUserCommandHandler(command LoginUserCommand, userRepository domainUsers.UserRepository) (string, error) {

	user, err := userRepository.GetVerifiedUser(command.Mail)
	if err != nil {
		return "", err
	}
	bycrypt := hashService.BCrypt{}
	if equal := bycrypt.Equal(command.Password, user.Password.Hashed); !equal {
		return "", ErrInvalidData
	}
	service := jwtService.New()
	token, err := service.GenerateToken(user.Pk.Val)
	if err != nil {
		return "", err
	}

	return token, nil

}
