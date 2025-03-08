package utils

import (
	model "ABTest/apps/Dao/userb/model"
	"ABTest/pkgs/proto/pb_userb"
	"errors"

	"google.golang.org/protobuf/proto"
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

func TranslateUserbModelToProtoUserbInfo(user *model.UserbBasic) (*pb_userb.UserInfo, error) {
	if user == nil {
		return nil, errEmptyModel
	}
	return &pb_userb.UserInfo{
		UserId:    proto.Uint64(user.UserID),
		Name:      proto.String(user.Name),
		Email:     proto.String(user.Email),
		Phone:     proto.String(user.Phone),
		Address:   proto.String(user.Address),
		Company:   proto.String(user.Company),
		Otherjson: proto.String(user.Otherjson),
	}, nil
}

func TranslateProtoUserbInfoToUserbModel(user *pb_userb.UserInfo) (*model.UserbBasic, error) {
	if user == nil {
		return nil, errEmptyProto
	}
	return &model.UserbBasic{
		UserID:    getUint64Value(user.UserId),
		Name:      getStringValue(user.Name),
		Email:     getStringValue(user.Email),
		Phone:     getStringValue(user.Phone),
		Address:   getStringValue(user.Address),
		Company:   getStringValue(user.Company),
		Otherjson: getStringValue(user.Otherjson),
	}, nil
}
