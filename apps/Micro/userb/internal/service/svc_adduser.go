package service

import (
	"context"

	"ABTest/apps/Dao/userb/model"
	conf "ABTest/apps/Micro/userb/internal/config"
	pb "ABTest/pkgs/proto/pb_userb"
	mysql "ABTest/pkgs/xmysql"
)

// implement the interface
func (s *userbService) AddUser(ctx context.Context, req *pb.AddUserRequest) (resp *pb.AddUserResponse, err error) {
	// ...existing code...
	conf.Log.Infof("AddUser begin: %v", req)
	resp = &pb.AddUserResponse{Success: false}
	db := mysql.GetDB()
	if db == nil {
		conf.Log.Errorf("AddUser get db failed")
		resp.Error = "get db failed"
		return resp, mysql.ERR_GetDB_FAILED
	}
	// 创建用户基本信息模型
	userModel := model.UserbBasic{
		UserID:  req.GetUserInfo().GetUserId(),
		Name:    req.GetUserInfo().GetName(),
		Email:   req.GetUserInfo().GetEmail(),
		Phone:   req.GetUserInfo().GetPhone(),
		Address: req.GetUserInfo().GetAddress(),
		Company: req.GetUserInfo().GetCompany(),
	}
	if err = db.Create(&userModel).Error; err != nil {
		conf.Log.Errorf("AddUser failed: %v", err)
		resp.Error = err.Error()
		return resp, err
	}
	resp.Success = true
	resp.UserId = userModel.UserID
	// ...existing code...
	conf.Log.Infof("AddUser end: %v", resp)
	return resp, nil
}
