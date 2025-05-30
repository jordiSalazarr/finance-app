package mysqlmembers

import (
	"finances.jordis.golang/infrastructure/dbmodels"
	"gorm.io/gorm"
)

type UsersGroupRepoMySQL struct {
	DB *gorm.DB
}

func NewUsersGroupRepoMySQL(db *gorm.DB) *UsersGroupRepoMySQL {
	return &UsersGroupRepoMySQL{
		DB: db,
	}
}
func (ug *UsersGroupRepoMySQL) Join(userId string, groupId string) error {
	err := ug.DB.Exec("INSERT INTO user_groups (user_pk, group_pk) VALUES (?, ?)", userId, groupId).Error
	if err != nil {
		return err
	}
	return nil
}
func (ug *UsersGroupRepoMySQL) GetUserGroups(userID string) ([]dbmodels.Group, error) {
	var groups []dbmodels.Group

	err := ug.DB.
		Table("groups").
		Joins("JOIN user_groups ON user_groups.group_pk = groups.pk").
		Where("user_groups.user_pk = ?", userID).
		Find(&groups).Error

	if err != nil {
		return nil, err
	}

	return groups, nil
}

func (ug *UsersGroupRepoMySQL) GetUsersFromGroup(groupId string) ([]string, error) {
	var userIds []string
	err := ug.DB.Table("user_groups").Select("user_pk").Where("group_pk = ?", groupId).Find(&userIds).Error
	if err != nil {
		return nil, err
	}
	return userIds, nil
}
