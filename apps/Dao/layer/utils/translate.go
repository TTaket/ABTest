package utils

import (
	model "ABTest/apps/Dao/layer/model"
	"ABTest/pkgs/proto/pb_layer"
	"errors"
)

func getUint64Value(value *uint64) uint64 {
	if value != nil {
		return *value
	}
	return 0
}

func getStringValue(value *string) string {
	if value != nil {
		return *value
	}
	return ""
}

// model
// UserID    uint64 `gorm:"primaryKey;autoIncrement;not null"`
// Name      string
// Email     string
// Phone     string
// Address   string
// Company   string
// Otherjson string

// proto
// UserID    uint64
// Name      string
// Email     string
// Phone     string
// Address   string
// Company   string
// Otherjson string
var (
	// empty model error
	errEmptyModel = errors.New("empty model")

	// empty proto error
	errEmptyProto = errors.New("empty proto")
)

func TranslateLayerBasicToProto(model *model.LayerBasic) (*pb_layer.Layer, error) {
	if model == nil {
		return nil, errEmptyModel
	}

	proto := &pb_layer.Layer{
		Id:          model.ID,
		Type:        pb_layer.LayerType(model.Type),
		Name:        model.Name,
		Description: model.Description,
		ParentId:    model.ParentID,
		Ratio:       model.Ratio,
	}

	return proto, nil
}

func TranslateLayerProtoToModel(proto *pb_layer.Layer) (*model.LayerBasic, error) {
	if proto == nil {
		return nil, errEmptyProto
	}

	model := &model.LayerBasic{
		ID:          proto.Id,
		Type:        int32(proto.Type),
		Name:        proto.Name,
		Description: proto.Description,
		ParentID:    proto.ParentId,
		Ratio:       proto.Ratio,
	}

	return model, nil
}
