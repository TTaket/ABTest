package model

import "time"

// LayerUserBinding 表示层与用户包的绑定关系
type LayerUserBinding struct {
	LayerID       string    `gorm:"column:layer_id;type:varchar(255)"`
	UserPackageID string    `gorm:"column:user_package_id;type:varchar(255)"`
	CreatedAt     time.Time // 创建时间
	UpdatedAt     time.Time // 更新时间
}
