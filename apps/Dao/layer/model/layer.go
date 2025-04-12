package model

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

type LayerBasic struct {
	ID          string  `gorm:"primaryKey;column:id;type:varchar(64);not null"`
	Type        int32   `gorm:"column:type;not null"`
	Name        string  `gorm:"column:name;type:varchar(255);not null"`
	Description string  `gorm:"column:description;type:varchar(255)"`
	ParentID    string  `gorm:"column:parent_id;type:varchar(64)"`
	Ratio       float32 `gorm:"column:ratio"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

func (table *LayerBasic) TableName() string {
	return "layer"
}

func (table *LayerBasic) String() string {
	return fmt.Sprintf("LayerBasic{ID: %s, Type: %d, Name: %s, Description: %s, ParentID: %s, Ratio: %.2f, CreatedAt: %s, UpdatedAt: %s, DeletedAt: %s}", table.ID, table.Type, table.Name, table.Description, table.ParentID, table.Ratio, table.CreatedAt.Format("2006-01-02 15:04:05"), table.UpdatedAt.Format("2006-01-02 15:04:05"), table.DeletedAt.Time.Format("2006-01-02 15:04:05"))
}
