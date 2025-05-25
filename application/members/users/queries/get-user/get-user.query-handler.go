package getuser

import (
	domainUsers "finances.jordis.golang/domain/members/users"
)

func GetUserQueyHandler(query GetUserQuery, usersRepo domainUsers.UserRepository) (domainUsers.User, error) {
	return usersRepo.GetById(query.UserID)
}
