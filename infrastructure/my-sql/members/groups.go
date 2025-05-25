package mysqlmembers

import (
	domainGroups "finances.jordis.golang/domain/members/groups"
	"finances.jordis.golang/infrastructure/dbmodels"
	"gorm.io/gorm"
)

type GroupsRepoMySQL struct {
	DB *gorm.DB
}

func NewGroupsRepoMySQL(db *gorm.DB) *GroupsRepoMySQL {
	return &GroupsRepoMySQL{
		DB: db,
	}
}

func (g *GroupsRepoMySQL) Save(Group domainGroups.Group) error {
	dbGroup := convertToDBGroup(Group)
	return g.DB.Create(&dbGroup).Error
}

func (g *GroupsRepoMySQL) Exists(name string) bool {
	var dbGroup dbmodels.Group
	err := g.DB.Model(&dbmodels.Group{}).Where("name = ?", name).First(&dbGroup).Error
	return err == nil
}
func (g *GroupsRepoMySQL) GetGroupByName(name string) (domainGroups.Group, error) {
	return domainGroups.Group{}, nil
}
func (g *GroupsRepoMySQL) CountUserCreatedGroups(userid string) int {
	return 0
}
func (g *GroupsRepoMySQL) GetById(id string) (domainGroups.Group, error) {
	var dbGroup dbmodels.Group
	err := g.DB.Model(&dbmodels.Group{}).Where("pk = ?", id).First(&dbGroup).Error
	if err != nil {
		return domainGroups.Group{}, err
	}

	return convertToDomainGroup(dbGroup)
}
func (g *GroupsRepoMySQL) GetBySecret(secret string) (domainGroups.Group, error) {
	var dbGroup dbmodels.Group
	err := g.DB.Model(&dbmodels.Group{}).Where("secret = ?", secret).First(&dbGroup).Error
	if err != nil {
		return domainGroups.Group{}, err
	}

	return convertToDomainGroup(dbGroup)

}

func convertManyToDomainGroup(dbGroups []dbmodels.Group) ([]domainGroups.Group, error) {
	var groups []domainGroups.Group
	for _, dbGroup := range dbGroups {
		group, err := convertToDomainGroup(dbGroup)
		if err != nil {
			return nil, err
		}
		groups = append(groups, group)
	}
	return groups, nil

}

func convertToDomainGroup(dbGroup dbmodels.Group) (domainGroups.Group, error) {
	return domainGroups.FromExistingGroup(dbGroup.PK, dbGroup.Name, dbGroup.Secret, dbGroup.CreatedBY, dbGroup.CreatedAt, dbGroup.UpdatedAt)
}

func convertToDBGroup(group domainGroups.Group) dbmodels.Group {
	return dbmodels.Group{
		PK:        group.Pk.Val,
		Name:      group.Name.Val,
		Secret:    group.Secret.Val,
		CreatedBY: group.Created_by.Val,
		CreatedAt: group.CreatedAt,
		UpdatedAt: group.UpdatedAt,
	}
}
