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
	return fmt.Sprintf("UserbBasic{UserID: %d, Name: %s, Email: %s, Phone: %s, Address: %s, Company: %s, Otherjson: %s, CreatedAt: %s, UpdatedAt: %s, DeletedAt: %s}", table.UserID, table.Name, table.Email, table.Phone, table.Address, table.Company, table.Otherjson, table.CreatedAt.Format("2006-01-02 15:04:05"), table.UpdatedAt.Format("2006-01-02 15:04:05"), table.DeletedAt.Time.Format("2006-01-02 15:04:05"))
}
