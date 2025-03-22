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
	return fmt.Sprintf("ExperimentGroupBasic{GroupID: %d, Name: %s, Description: %s, Allocation: %f, FromExperimentID: %d, WhiteListUserPackageId: %d, CreatedAt: %s, UpdatedAt: %s, DeletedAt: %s}", table.GroupID, table.Name, table.Description, table.Allocation, table.FromExperimentID, table.WhiteListUserPackageId, table.CreatedAt.Format("2006-01-02 15:04:05"), table.UpdatedAt.Format("2006-01-02 15:04:05"), table.DeletedAt.Time.Format("2006-01-02 15:04:05"))
}
