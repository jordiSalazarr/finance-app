package getusergroups

import (
	"fmt"

	"finances.jordis.golang/domain/members"
	domainGroups "finances.jordis.golang/domain/members/groups"
)

func GetUserGroupsQueryHandler(query GetUserGroupsQuery, groupRepo domainGroups.GroupRepository, usersGroupsRepository members.UsersGroupRepository) ([]*domainGroups.Group, error) {
	groupsIds, err := usersGroupsRepository.GetUserGroups(query.UserId)
	if err != nil {
		return nil, err
	}

	var groups []*domainGroups.Group
	for _, id := range groupsIds {
		group, err := groupRepo.GetById(id)
		if err != nil {
			fmt.Print(err.Error())
			continue
		}
		groups = append(groups, &group)
	}

	return groups, nil
}
