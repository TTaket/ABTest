package model

import (
	"fmt"

	"gorm.io/gorm"
)

// UserbModel UserbModel
type UserbBasic struct {
	gorm.Model

	UserID    uint64 `gorm:"primaryKey;autoIncrement;not null"`
	Name      string
	Email     string
	Phone     string
	Address   string
	Company   string
	Otherjson string
}

func (table *UserbBasic) TableName() string {
	return "userb"
}

func (table *UserbBasic) String() string {
	return "user_id: " + fmt.Sprint(table.UserID) + ", name: " + table.Name + ", email: " + table.Email + ", phone: " + table.Phone + ", address: " + table.Address + ", company: " + table.Company + ", otherjson: " + table.Otherjson
}
