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
	return "experiment_id: " + fmt.Sprint(table.ExperimentID) + ", name: " + table.Name + ", description: " + table.Description + ", status: " + fmt.Sprint(table.Status)
}
