package model

import (
	"errors"

	gorm "gorm.io/gorm"
)

var (
	errExperimentIDRequired      = errors.New("experiment_id is required but not found")
	errExperimentGroupIDRequired = errors.New("group_id is required but not found")
)

// experimentinfo
func (table *ExperimentInfoBasic) CreateExperimentInfo(db *gorm.DB) error {
	err := db.Create(table).Error
	if err != nil {
		return err
	}
	return nil
}
func (table *ExperimentInfoBasic) UpdateExperimentInfo(db *gorm.DB) error {
	if table.ExperimentID == 0 {
		return errExperimentIDRequired
	}
	return db.Model(table).Updates(table).Error
}
func (table *ExperimentInfoBasic) DeleteExperimentInfo(db *gorm.DB) error {
	if table.ExperimentID == 0 {
		return errExperimentIDRequired
	}
	return db.Delete(table).Error
}

// 根据id操作
func DeleteExperimentInfoByID(db *gorm.DB, experimentID uint64) error {
	if experimentID == 0 {
		return errExperimentIDRequired
	}
	return db.Where("experiment_id = ?", experimentID).Delete(&ExperimentInfoBasic{}).Error
}
func GetExperimentInfoByID(db *gorm.DB, experimentID uint64) (*ExperimentInfoBasic, error) {
	if experimentID == 0 {
		return nil, errExperimentIDRequired
	}
	experimentInfo := &ExperimentInfoBasic{}
	err := db.Where("experiment_id = ?", experimentID).First(experimentInfo).Error
	return experimentInfo, err
}

// experimentgroup
func (table *ExperimentGroupBasic) CreateExperimentGroup(db *gorm.DB) error {
	err := db.Create(table).Error
	if err != nil {
		return err
	}
	return nil
}
func (table *ExperimentGroupBasic) UpdateExperimentGroup(db *gorm.DB) error {
	if table.GroupID == 0 {
		return errExperimentGroupIDRequired
	}
	// 部分更新
	return db.Model(table).Updates(table).Error
}
func (table *ExperimentGroupBasic) DeleteExperimentGroup(db *gorm.DB) error {
	if table.GroupID == 0 {
		return errExperimentGroupIDRequired
	}
	return db.Delete(table).Error
}

// 根据id操作
func DeleteExperimentGroupByID(db *gorm.DB, groupID uint64) error {
	if groupID == 0 {
		return errExperimentGroupIDRequired
	}
	return db.Where("group_id = ?", groupID).Delete(&ExperimentGroupBasic{}).Error
}
func GetExperimentGroupByID(db *gorm.DB, groupID uint64) (*ExperimentGroupBasic, error) {
	if groupID == 0 {
		return nil, errExperimentGroupIDRequired
	}
	experimentGroup := &ExperimentGroupBasic{}
	err := db.Where("group_id = ?", groupID).First(experimentGroup).Error
	return experimentGroup, err
}

// 获取某一个实验的所有分组
func GetExperimentGroupsByExperimentID(db *gorm.DB, experimentID uint64) ([]ExperimentGroupBasic, error) {
	experimentGroups := []ExperimentGroupBasic{}
	err := db.Where("from_experiment_id = ?", experimentID).Find(&experimentGroups).Error
	return experimentGroups, err
}

// 根据实验id删除所有实验组
func DeleteExperimentGroupsByExperimentID(db *gorm.DB, experimentID uint64) error {
	if experimentID == 0 {
		return errExperimentIDRequired
	}
	return db.Where("from_experiment_id = ?", experimentID).Delete(&ExperimentGroupBasic{}).Error
}
