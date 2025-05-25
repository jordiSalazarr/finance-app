package dbmodels

import (
	"time"
)

type Transaction struct {
	PK           string `gorm:"primaryKey"`
	Description  string
	Amount       int64 // representa los valores en centavos, o usa una lib como shopspring/decimal
	Type         string
	Category     string
	AlreadyPayed bool
	PayedBy      string
	GroupPK      string
	UserPK       string

	CreatedAt time.Time
	UpdatedAt time.Time
}
