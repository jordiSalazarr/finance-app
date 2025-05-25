package domainGroups

type GroupRepository interface {
	Save(Group Group) error
	Exists(name string) bool
	GetGroupByName(name string) (Group, error)
	CountUserCreatedGroups(userid string) int
	GetById(id string) (Group, error)
	GetBySecret(secret string) (Group, error)
}
