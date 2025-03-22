package model

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

// ExperimentModel ExperimentInfo
type ExperimentInfoBasic struct {
	ExperimentID uint64 `gorm:"primaryKey;column:experiment_id;autoIncrement;not null"`
	Name         string `gorm:"column:name;type:varchar(255)"`
	Description  string `gorm:"column:description;type:varchar(255)"`
	Status       int    `gorm:"column:status;type:int"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (table *ExperimentInfoBasic) TableName() string {
	return "experimentinfo"
}

func (table *ExperimentInfoBasic) String() string {
	return fmt.Sprintf("ExperimentInfoBasic{ExperimentID: %d, Name: %s, Description: %s, Status: %d, CreatedAt: %s, UpdatedAt: %s, DeletedAt: %s}", table.ExperimentID, table.Name, table.Description, table.Status, table.CreatedAt.Format("2006-01-02 15:04:05"), table.UpdatedAt.Format("2006-01-02 15:04:05"), table.DeletedAt.Time.Format("2006-01-02 15:04:05"))
}
