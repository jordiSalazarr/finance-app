package members

import "finances.jordis.golang/infrastructure/dbmodels"

type UsersGroupRepository interface {
	Join(userId string, groupId string) error
	GetUserGroups(userID string) ([]dbmodels.Group, error)
	GetUsersFromGroup(groupId string) ([]string, error)
}
