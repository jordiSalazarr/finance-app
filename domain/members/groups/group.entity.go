package domainGroups

import (
	"time"

	"finances.jordis.golang/domain"
	groupsVals "finances.jordis.golang/domain/members/groups/value-objects"
)

type Group struct {
	Pk         domain.UUID
	Name       groupsVals.Name
	Secret     groupsVals.Secret
	Created_by domain.UUID
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

func NewGroup(nameStr string, createdBy domain.UUID) (Group, error) {
	name, err := groupsVals.NewGroupName(nameStr)
	if err != nil {
		return Group{}, err
	}

	secret := groupsVals.NewSecret()
	pk := domain.UUID{Val: domain.NewUUID()}
	createdAt, updatedAt := time.Now(), time.Now()

	return Group{
		Pk:         pk,
		Name:       name,
		Secret:     secret,
		Created_by: createdBy,
		CreatedAt:  createdAt,
		UpdatedAt:  updatedAt,
	}, nil

}

func FromExistingGroup(
	pk string,
	name string,
	secret string,
	createdBy string,
	createdAt time.Time,
	updatedAt time.Time,
) (Group, error) {
	id := domain.UUID{Val: pk}
	createdByUUID := domain.UUID{Val: createdBy}
	nameVal, err := groupsVals.NewGroupName(name)
	if err != nil {
		return Group{}, err
	}

	secretVal := groupsVals.ExistingSecret(secret)

	return Group{
		Pk:         id,
		Name:       nameVal,
		Secret:     secretVal,
		Created_by: createdByUUID,
		CreatedAt:  createdAt,
		UpdatedAt:  updatedAt,
	}, nil
}
