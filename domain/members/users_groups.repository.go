package members

type UsersGroupRepository interface {
	Join(userId string, groupId string) error
	GetUserGroups(userID string) ([]string, error)
	GetUsersFromGroup(groupId string) ([]string, error)
}
