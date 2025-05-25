package usersgorupsInMemory

import (
	"errors"

	domainUsers "finances.jordis.golang/domain/members/users"
)

var (
	ErrGroupNotFound = errors.New("group does not exist")
)

type UsersgorupsInMemory struct {
	UsersGroups map[string][]string
}

func New() *UsersgorupsInMemory {
	return &UsersgorupsInMemory{
		UsersGroups: map[string][]string{},
	}
}
func (im *UsersgorupsInMemory) Join(userId string, groupId string) error {
	if _, exists := im.UsersGroups[groupId]; !exists {
		im.UsersGroups[groupId] = []string{userId}
		return nil
	}

	im.UsersGroups[groupId] = append(im.UsersGroups[groupId], userId)
	return nil

}

func (im *UsersgorupsInMemory) GetUserGroups(userID string) ([]string, error) {
	groups := []string{}
	for groupId, usersIds := range im.UsersGroups {
		for _, id := range usersIds {
			if id == userID {
				groups = append(groups, groupId)
			}
		}
	}

	return groups, nil
}

func (im *UsersgorupsInMemory) GetUsersFromGroup(groupId string) ([]string, error) {
	if users, ok := im.UsersGroups[groupId]; ok {
		return users, nil
	}

	return nil, ErrGroupNotFound

}

func (im *UsersgorupsInMemory) Exists(id string) bool {
	_, ok := im.UsersGroups[id]
	return ok
}

func (im *UsersgorupsInMemory) GetAll() ([]domainUsers.User, error) {
	return []domainUsers.User{}, nil
}

// func (im *UsersgorupsInMemory) GetById(id string) bool {
// 	user, ok := im.UsersGroups[id]
// 	return user
// }
