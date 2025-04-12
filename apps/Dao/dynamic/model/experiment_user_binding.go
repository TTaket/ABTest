package model

import "time"

// ExperimentUserBinding 表示实验组与用户包的绑定关系
type ExperimentUserBinding struct {
	ExperimentGroupID string    `gorm:"column:experiment_group_id;type:varchar(255)"`
	UserPackageID     string    `gorm:"column:user_package_id;type:varchar(255)"`
	CreatedAt         time.Time // 创建时间
	UpdatedAt         time.Time // 更新时间
}
