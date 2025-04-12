package model

import "time"

// ExperimentLayerBinding 表示实验与层的绑定关系
type ExperimentLayerBinding struct {
	ExperimentID string    `gorm:"column:experiment_id;type:varchar(255)"`
	LayerID      string    `gorm:"column:layer_id;type:varchar(255)"`
	Ratio        float32   `gorm:"column:ratio;type:float"`
	CreatedAt    time.Time // 创建时间
	UpdatedAt    time.Time // 更新时间
}
