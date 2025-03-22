package model

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

type ExperimentGroupBasic struct {
	GroupID                uint64  `gorm:"primaryKey;column:group_id;autoIncrement;not null"`
	Name                   string  `gorm:"column:name;type:varchar(255)"`
	Description            string  `gorm:"column:description;type:varchar(255)"`
	Allocation             float32 `gorm:"column:allocation;type:float"`
	FromExperimentID       uint64  `gorm:"column:from_experiment_id;"`
	WhiteListUserPackageId uint64  `gorm:"column:white_list_user_package_id;"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (table *ExperimentGroupBasic) TableName() string {
	return "experimentgroup"
}

func (table *ExperimentGroupBasic) String() string {
	return "group_id: " + fmt.Sprint(table.GroupID) + ", name: " + table.Name + ", description: " + table.Description + ", allocation: " + fmt.Sprint(table.Allocation) + ", from_experiment_id: " + fmt.Sprint(table.FromExperimentID)
}
