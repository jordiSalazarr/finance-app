package groups_useCases

import (
	"errors"

	joingroup "finances.jordis.golang/application/members/users_groups/commands/join-group"
	"finances.jordis.golang/domain"
	"finances.jordis.golang/domain/members"
	domainGroups "finances.jordis.golang/domain/members/groups"
)

const maxGroupsCreatedByUser = 10

var (
	ErrMaxGroupsReached = errors.New("a user can only create a max of 10 groups")
)

func CreateGroupcommandHandler(command CreateGroupCommand, groupsRepository domainGroups.GroupRepository, usersGroupsRepo members.UsersGroupRepository) (domainGroups.Group, error) {

	if currCount := groupsRepository.CountUserCreatedGroups(command.CreatedBy); currCount >= maxGroupsCreatedByUser {
		return domainGroups.Group{}, ErrMaxGroupsReached
	}
	user_id := domain.UUID{
		Val: command.CreatedBy,
	}
	group, err := domainGroups.NewGroup(command.Name, user_id)
	if err != nil {
		return domainGroups.Group{}, err
	}

	err = groupsRepository.Save(group)
	if err != nil {
		return domainGroups.Group{}, err
	}
	joinGroupcommand := joingroup.JoinGroupCommand{
		Secret: group.Secret.Val,
		UserId: command.CreatedBy,
	}
	err = joingroup.JoinGroupCommandHandler(joinGroupcommand, usersGroupsRepo, groupsRepository)
	return group, err

}
