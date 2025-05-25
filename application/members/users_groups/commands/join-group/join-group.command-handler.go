package joingroup

import (
	"errors"

	"finances.jordis.golang/domain/members"
	domainGroups "finances.jordis.golang/domain/members/groups"
)

var (
	ErrInvalidSecret = errors.New("invalid secret")
)

func JoinGroupCommandHandler(command JoinGroupCommand, users_gorups_repository members.UsersGroupRepository, groupsRepo domainGroups.GroupRepository) error {
	group, err := groupsRepo.GetBySecret(command.Secret)
	if err != nil {
		return err
	}
	if group.Secret.Val != command.Secret {
		return ErrInvalidSecret
	}
	return users_gorups_repository.Join(command.UserId, group.Pk.Val)
}
