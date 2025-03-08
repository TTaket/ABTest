package model

import (
	"errors"

	"gorm.io/gorm"
)

var (
	// err userid is required but not found
	errUserIDRequired = errors.New("userid is required but not found")
)

// 增删改查
func (table *UserbBasic) CreateUserb(db *gorm.DB) error {
	err := db.Create(table).Error
	if err != nil {
		return err
	}
	table.UserID = uint64(table.ID)
	return nil
}

func (table *UserbBasic) UpdateUserb(db *gorm.DB) error {
	if table.UserID == 0 {
		return errUserIDRequired
	}
	return db.Save(table).Error
}

func (table *UserbBasic) DeleteUserb(db *gorm.DB) error {
	if table.UserID == 0 {
		return errUserIDRequired
	}
	return db.Delete(table).Error
}
func DeleteUserbByID(db *gorm.DB, userID uint64) error {
	if userID == 0 {
		return errUserIDRequired
	}
	return db.Where("user_id = ?", userID).Delete(&UserbBasic{}).Error
}

func GetUserbByID(db *gorm.DB, userID uint64) (*UserbBasic, error) {
	if userID == 0 {
		return nil, errUserIDRequired
	}
	userb := &UserbBasic{}
	err := db.Where("user_id = ?", userID).First(userb).Error
	return userb, err
}
