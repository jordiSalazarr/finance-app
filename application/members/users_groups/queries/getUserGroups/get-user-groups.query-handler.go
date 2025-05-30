package getusergroups

import (
	"fmt"

	"finances.jordis.golang/domain/members"
	domainGroups "finances.jordis.golang/domain/members/groups"
)

func GetUserGroupsQueryHandler(query GetUserGroupsQuery, groupRepo domainGroups.GroupRepository, usersGroupsRepository members.UsersGroupRepository) ([]domainGroups.Group, error) {
	dbgroups, err := usersGroupsRepository.GetUserGroups(query.UserId)
	if err != nil {
		return nil, err
	}
	if len(dbgroups) == 0 {
		return nil, nil
	}

	groups := make([]domainGroups.Group, len(dbgroups))
	for i, group := range dbgroups {
		domainGroup, err := domainGroups.FromExistingGroup(
			group.PK,
			group.Name,
			group.Secret,
			group.CreatedBY,
			group.CreatedAt,
			group.UpdatedAt,
		)
		if err != nil {
			fmt.Printf("Error creating domain group from existing group: %v\n", err)
		} else {
			groups[i] = domainGroup
		}

	}

	return groups, nil
}
