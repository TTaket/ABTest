package model

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

// UserbModel UserbModel
type UserbBasic struct {
	UserID    uint64 `gorm:"primaryKey;column:user_id;autoIncrement;not null"`
	Name      string `gorm:"column:name;type:varchar(64)"`
	Email     string `gorm:"column:email;type:varchar(64)"`
	Phone     string `gorm:"column:phone;type:varchar(32)"`
	Address   string `gorm:"column:address;type:varchar(255)"`
	Company   string `gorm:"column:company;type:varchar(255)"`
	Otherjson string `gorm:"column:otherjson;type:json"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (table *UserbBasic) TableName() string {
	return "userb"
}

func (table *UserbBasic) String() string {
	return "user_id: " + fmt.Sprint(table.UserID) + ", name: " + table.Name + ", email: " + table.Email + ", phone: " + table.Phone + ", address: " + table.Address + ", company: " + table.Company + ", otherjson: " + table.Otherjson
}
