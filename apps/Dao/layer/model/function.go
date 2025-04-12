package model

import (
	"errors"

	"gorm.io/gorm"
)

var (
	errLayerIDRequired = errors.New("layer ID is required but not found")
)

func (table *LayerBasic) CreateLayer(db *gorm.DB) error {
	return db.Create(table).Error
}

func (table *LayerBasic) UpdateLayer(db *gorm.DB) error {
	if table.ID == "" {
		return errLayerIDRequired
	}
	return db.Model(table).Updates(table).Error
}

func (table *LayerBasic) DeleteLayer(db *gorm.DB) error {
	if table.ID == "" {
		return errLayerIDRequired
	}
	return db.Delete(table).Error
}

func DeleteLayerByID(db *gorm.DB, id string) error {
	if id == "" {
		return errLayerIDRequired
	}
	return db.Where("id = ?", id).Delete(&LayerBasic{}).Error
}

func GetLayerByID(db *gorm.DB, id string) (*LayerBasic, error) {
	if id == "" {
		return nil, errLayerIDRequired
	}
	layer := &LayerBasic{}
	err := db.Where("id = ?", id).First(layer).Error
	return layer, err
}

func ListLayers(db *gorm.DB, layerType int32, parentID string, page, pageSize int) ([]LayerBasic, int64, error) {
	var layers []LayerBasic
	var total int64

	query := db.Model(&LayerBasic{})
	if layerType >= 0 {
		query = query.Where("type = ?", layerType)
	}
	if parentID != "" {
		query = query.Where("parent_id = ?", parentID)
	}

	err := query.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * pageSize
	err = query.Offset(offset).Limit(pageSize).Find(&layers).Error
	return layers, total, err
}
