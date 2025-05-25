package inMemoryUsers

import (
	"errors"

	domainUsers "finances.jordis.golang/domain/members/users"
)

var (
	ErrUserNotVerified = errors.New("user is not verified yet")
	ErrUserNotFound    = errors.New("user not found")
)

type InMemoryUsersRepo struct {
	Users []domainUsers.User
}

func New() *InMemoryUsersRepo {
	return &InMemoryUsersRepo{
		Users: []domainUsers.User{},
	}
}

func (im *InMemoryUsersRepo) Save(user domainUsers.User) error {
	im.Users = append(im.Users, user)
	return nil
}
func (im *InMemoryUsersRepo) Exists(mail string) bool {
	for _, user := range im.Users {
		if user.Mail.Val == mail {
			return true
		}
	}

	return false
}

func (im *InMemoryUsersRepo) GetVerifiedUser(mail string) (domainUsers.User, error) {
	var dbUser domainUsers.User
	for _, user := range im.Users {
		if user.Mail.Val == mail {
			if !user.IsActive || !user.IsVerified {
				return domainUsers.User{}, ErrUserNotVerified
			}

			dbUser = user
			return dbUser, nil
		}
	}

	return domainUsers.User{}, ErrUserNotFound

}

func (im *InMemoryUsersRepo) VerificateUser(mail string) error {
	for idx, user := range im.Users {
		if user.Mail.Val == mail {
			user.IsActive = true
			user.IsVerified = true
			im.Users[idx] = user
			return nil
		}
	}

	return ErrUserNotFound
}

func (im *InMemoryUsersRepo) GetUser(mail string) (domainUsers.User, error) {
	for _, user := range im.Users {
		if user.Mail.Val == mail {
			return user, nil
		}
	}

	return domainUsers.User{}, ErrUserNotFound
}

func (im *InMemoryUsersRepo) GetById(id string) (domainUsers.User, error) {
	for _, user := range im.Users {
		if user.Pk.Val == id {
			return user, nil
		}
	}

	return domainUsers.User{}, ErrUserNotFound
}

func (im *InMemoryUsersRepo) UpdateCurrentBalance(id string, val int64) error {
	for _, user := range im.Users {
		if user.Pk.Val == id {
			user.CurrentBalance.Val += val
			return nil
		}
	}

	return ErrUserNotFound
}

func (im *InMemoryUsersRepo) UpdateActorsCurrentBalance(debtor_id, payed_by string, val int64) error {
	for _, user := range im.Users {
		if user.Pk.Val == debtor_id {
			user.CurrentBalance.Val -= val
		}
		if user.Pk.Val == payed_by {
			user.CurrentBalance.Val += val
		}
	}

	return nil
}

func (im *InMemoryUsersRepo) GetAll() ([]domainUsers.User, error) {
	return []domainUsers.User{}, nil
}
