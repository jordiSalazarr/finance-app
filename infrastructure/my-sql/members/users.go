package mysqlmembers

import (
	"fmt"
	"log"
	"time"

	domainUsers "finances.jordis.golang/domain/members/users"
	"finances.jordis.golang/infrastructure/dbmodels"
	"gorm.io/gorm"
)

type UsersRepoMySQL struct {
	DB *gorm.DB
}

func NewUsersRepoMySQL(db *gorm.DB) *UsersRepoMySQL {
	return &UsersRepoMySQL{
		DB: db,
	}
}

func (u *UsersRepoMySQL) GetAll() ([]domainUsers.User, error) {
	var dbUsers []dbmodels.User
	err := u.DB.Find(&dbUsers).Error
	if err != nil {
		return nil, err
	}

	domainUsers := make([]domainUsers.User, len(dbUsers))
	for i, dbUser := range dbUsers {
		domainUser, err := mapToDomainUser(dbUser)
		if err != nil {
			return nil, err
		}
		domainUsers[i] = domainUser
	}
	return domainUsers, nil
}

func (u *UsersRepoMySQL) Save(user domainUsers.User) error {
	dbUser := mapDomainToDBUser(user)

	err := u.DB.Create(&dbUser).Error
	if err != nil {
		log.Printf("❌ error creating user: %v", err)
		return fmt.Errorf("failed to save user: %w", err)
	}
	return nil
}

func (u *UsersRepoMySQL) Exists(mail string) bool {
	var dbUser dbmodels.User
	err := u.DB.Model(&dbmodels.User{}).Where("mail = ?", mail).First(&dbUser).Error
	return err == nil
}

func (u *UsersRepoMySQL) GetVerifiedUser(mail string) (domainUsers.User, error) {
	var dbUser dbmodels.User
	err := u.DB.Where("mail = ? AND is_verified = TRUE", mail).First(&dbUser).Error
	if err != nil {
		return domainUsers.User{}, err
	}
	return mapToDomainUser(dbUser)
}

// TODO: implement this shit
func (u *UsersRepoMySQL) VerificateUser(mail string) error {
	return nil
}

func (u *UsersRepoMySQL) GetUser(mail string) (domainUsers.User, error) {
	var dbUser dbmodels.User
	err := u.DB.Where("mail = ?", mail).First(&dbUser).Error
	if err != nil {
		return domainUsers.User{}, err
	}

	return mapToDomainUser(dbUser)
}
func (u *UsersRepoMySQL) GetById(id string) (domainUsers.User, error) {
	var dbUser dbmodels.User
	err := u.DB.Where("pk = ?", id).First(&dbUser).Error
	if err != nil {
		return domainUsers.User{}, err
	}

	return mapToDomainUser(dbUser)
}

func (u *UsersRepoMySQL) UpdateCurrentBalance(id string, val int64) error {
	return updateUserCurrentBalance(id, val, u.DB)

}

func (u *UsersRepoMySQL) UpdateActorsCurrentBalance(debtorID, payedBy string, val int64) error {
	return u.DB.Transaction(func(tx *gorm.DB) error {
		if err := updateUserCurrentBalance(debtorID, -val, tx); err != nil {
			return err
		}
		if err := updateUserCurrentBalance(payedBy, val, tx); err != nil {
			return err
		}
		return nil
	})
}

func updateUserCurrentBalance(userID string, val int64, tx *gorm.DB) error {
	var user dbmodels.User
	if err := tx.Where("pk = ?", userID).First(&user).Error; err != nil {
		return err
	}

	newBalance := user.CurrentBalance + val
	if newBalance < 0 {
		return domainUsers.ErrNegativeBalanceError
	}
	return tx.Model(&dbmodels.User{}).Where("pk = ?", userID).Update("current_balance", newBalance).Error
}

func mapToDomainUser(dbUser dbmodels.User) (domainUsers.User, error) {
	return domainUsers.FromDBUser(dbUser.PK, dbUser.UserName, dbUser.MonthlyIncome, dbUser.Mail, dbUser.Password, dbUser.CurrentBalance, dbUser.VerificationCode, dbUser.VerificationCodeExpiryDate)

}

func mapDomainToDBUser(u domainUsers.User) dbmodels.User {
	return dbmodels.User{
		PK:                         u.Pk.Val,
		UserName:                   u.Username.Val,
		Mail:                       u.Mail.Val,
		Password:                   u.Password.Hashed,
		CurrentBalance:             (u.CurrentBalance.Val),
		MonthlyIncome:              (u.MonthlyIncome.Val),
		IsActive:                   u.IsActive,
		IsVerified:                 u.IsVerified,
		VerificationCode:           u.VerificationCode.Val,
		CreatedAt:                  u.CreatedAt,
		UpdatedAt:                  u.UpdatedAt,
		VerificationCodeExpiryDate: u.VerificationCodeExpiry.In(time.UTC),
	}
}
