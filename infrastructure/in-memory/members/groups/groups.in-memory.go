package inmemoryGroups

import (
	"errors"

	domainGroups "finances.jordis.golang/domain/members/groups"
)

var (
	ErrGroupNotFound = errors.New("group not found")
)

type InMemoryGroups struct {
	Groups []domainGroups.Group
}

func New() *InMemoryGroups {
	return &InMemoryGroups{
		Groups: []domainGroups.Group{},
	}
}

func (im *InMemoryGroups) Save(group domainGroups.Group) error {
	im.Groups = append(im.Groups, group)
	return nil
}

func (im *InMemoryGroups) Exists(name string) bool {
	return true
}

func (im *InMemoryGroups) CountUserCreatedGroups(mail string) int {
	total := 0
	for _, group := range im.Groups {
		if group.Created_by.Val == mail {
			total++
		}
	}

	return total
}

func (im *InMemoryGroups) GetGroupByName(name string) (domainGroups.Group, error) {
	for _, group := range im.Groups {
		if group.Name.Val == name {
			return group, nil
		}
	}

	return domainGroups.Group{}, ErrGroupNotFound
}

func (im *InMemoryGroups) GetById(id string) (domainGroups.Group, error) {
	for _, group := range im.Groups {
		if group.Pk.Val == id {
			return group, nil
		}
	}

	return domainGroups.Group{}, ErrGroupNotFound
}
