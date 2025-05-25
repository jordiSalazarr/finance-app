package mysqlmembers

import "gorm.io/gorm"

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
func (ug *UsersGroupRepoMySQL) GetUserGroups(userID string) ([]string, error) {
	var groupIds []string
	err := ug.DB.Table("user_groups").Select("group_pk").Where("user_pk = ?", userID).Find(&groupIds).Error
	if err != nil {
		return nil, err
	}
	return groupIds, nil
}
func (ug *UsersGroupRepoMySQL) GetUsersFromGroup(groupId string) ([]string, error) {
	var userIds []string
	err := ug.DB.Table("user_groups").Select("user_pk").Where("group_pk = ?", groupId).Find(&userIds).Error
	if err != nil {
		return nil, err
	}
	return userIds, nil
}
