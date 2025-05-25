package dbmodels

import "time"

type Group struct {
	PK        string `gorm:"primaryKey"`
	Name      string
	Secret    string
	CreatedBY string
	Users     []*User `gorm:"many2many:user_groups;"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
