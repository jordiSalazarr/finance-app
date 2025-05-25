package domainUsers

import (
	"time"

	"finances.jordis.golang/domain"
	userVals "finances.jordis.golang/domain/members/users/value-objects"
	hashService "finances.jordis.golang/services"
)

type User struct {
	Pk                     domain.UUID
	Username               userVals.Username
	Mail                   userVals.Mail
	Password               userVals.Password
	CurrentBalance         userVals.Balance
	MonthlyIncome          userVals.MonthlyIncome
	IsActive               bool
	IsVerified             bool
	VerificationCode       userVals.VerificationCode
	VerificationCodeExpiry time.Time
	CreatedAt              time.Time
	UpdatedAt              time.Time
}

func NewUser(usernameStr string, monthlyIncomeFlt int64, mailStr string, plainPassword string, currentBalanceInt int64, hashService *hashService.BCrypt) (User, error) {
	username, err := userVals.NewUsername(usernameStr)
	if err != nil {
		return User{}, err
	}

	monthlyIncome, err := userVals.NewMonthlyIncome(monthlyIncomeFlt)
	if err != nil {
		return User{}, err
	}
	mail, err := userVals.NewMail(mailStr)
	if err != nil {
		return User{}, err
	}

	password, err := userVals.NewPassword(plainPassword, hashService)
	if err != nil {
		return User{}, err
	}

	currentBalance := userVals.NewUserBalance(currentBalanceInt)
	pk := domain.UUID{
		Val: domain.NewUUID(),
	}
	code := userVals.NewVerificationCode()
	createdAt := time.Now()
	updatedAt := time.Now()
	verificationCodeExpiry := time.Now().Add(30 * time.Minute)
	return User{
		Pk:                     pk,
		Username:               username,
		Mail:                   mail,
		Password:               password,
		CurrentBalance:         currentBalance,
		MonthlyIncome:          monthlyIncome,
		IsActive:               false,
		IsVerified:             false,
		VerificationCode:       code,
		CreatedAt:              createdAt,
		UpdatedAt:              updatedAt,
		VerificationCodeExpiry: verificationCodeExpiry,
	}, nil

}

func FromDBUser(pk string, usernameStr string, monthlyIncomeFlt int64, mailStr string, hashedPassword string, currentBalanceInt int64, verificationCode string, verificationCodeExpiry time.Time) (User, error) {

	domainPK := domain.UUID{Val: pk}

	username, err := userVals.NewUsername(usernameStr)
	if err != nil {
		return User{}, err
	}

	monthlyIncome, err := userVals.NewMonthlyIncome(monthlyIncomeFlt)
	if err != nil {
		return User{}, err
	}
	mail, err := userVals.NewMail(mailStr)
	if err != nil {
		return User{}, err
	}
	password := userVals.Password{
		Hashed: hashedPassword,
	}

	currentBalance := userVals.NewUserBalance(currentBalanceInt)

	code := userVals.ExistingVerificationCode(verificationCode)
	createdAt := time.Now()
	updatedAt := time.Now()
	return User{
		Pk:                     domainPK,
		Username:               username,
		Mail:                   mail,
		Password:               password,
		CurrentBalance:         currentBalance,
		MonthlyIncome:          monthlyIncome,
		IsActive:               false,
		IsVerified:             false,
		VerificationCode:       code,
		CreatedAt:              createdAt,
		UpdatedAt:              updatedAt,
		VerificationCodeExpiry: verificationCodeExpiry,
	}, nil
}
