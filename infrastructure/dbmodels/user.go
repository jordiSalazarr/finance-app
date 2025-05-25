package dbmodels

import (
	"time"
)

type User struct {
	PK                         string `gorm:"primaryKey"`
	UserName                   string
	Mail                       string
	Password                   string
	CurrentBalance             int64
	MonthlyIncome              int64
	IsActive                   bool
	IsVerified                 bool
	VerificationCode           string
	VerificationCodeExpiryDate time.Time
	CreatedAt                  time.Time
	UpdatedAt                  time.Time
	Groups                     []*Group `gorm:"many2many:user_groups;"`
	Transactions               []Transaction
}
