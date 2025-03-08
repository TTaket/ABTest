package service

import (
	"context"
	"errors"

	model "ABTest/apps/Dao/userb/model"
	modelutils "ABTest/apps/Dao/userb/utils"
	pb "ABTest/pkgs/proto/pb_userb"
	"ABTest/pkgs/xmysql"
)

var (
	// error - 请求体没有id
	errIdRequired = errors.New("id is required but not found")

	// err - db获取失败
	errGetDBFailed = xmysql.ERR_GetDB_FAILED

	// err - 整体失败
	errFailed = errors.New("failed")
)

func (s *userbService) GetUserInfo(ctx context.Context, in *pb.GetUserInfoRequest) (*pb.GetUserInfoResponse, error) {
	// 创建返回体
	out := &pb.GetUserInfoResponse{
		Success: false,
	}
	// 获得db
	db := xmysql.GetDB()
	if db == nil {
		out.Error = errGetDBFailed.Error()
		return out, errGetDBFailed
	}

	// 查询用户
	userbmodel, err := model.GetUserbByID(db, *in.UserInfo.UserId)
	if err != nil {
		out.Error = err.Error()
		return out, err
	}
	// 返回结果
	userinfo, err := modelutils.TranslateUserbModelToProtoUserbInfo(userbmodel)
	if err != nil {
		out.Error = err.Error()
		return out, err
	}

	out.Success = true
	out.UserInfo = userinfo
	return out, nil
}

func (s *userbService) BatchGetUserInfo(ctx context.Context, in *pb.BatchGetUserInfoRequest) (*pb.BatchGetUserInfoResponse, error) {
	// 创建返回体
	out := &pb.BatchGetUserInfoResponse{
		Success: false,
	}
	var errflag bool = false
	var userinfos []*pb.UserInfo
	var errinfo []string
	// 获得db
	db := xmysql.GetDB()
	if db == nil {
		out.Errors = append(out.Errors, errGetDBFailed.Error())
		return out, errGetDBFailed
	}
	// 取出id
	if in.UserInfo == nil {
		out.Errors = append(out.Errors, errIdRequired.Error())
		return out, errIdRequired
	}
	// 查询用户
	for _, userinfo := range in.UserInfo {
		userbmodels, err := model.GetUserbByID(db, *userinfo.UserId)
		if err != nil {
			errflag = true
			errinfo = append(errinfo, errIdRequired.Error())
			userinfos = append(userinfos, nil)
			continue
		}

		retuserinfo, err := modelutils.TranslateUserbModelToProtoUserbInfo(userbmodels)
		if err != nil {
			errflag = true
			errinfo = append(errinfo, errIdRequired.Error())
			userinfos = append(userinfos, nil)
			continue
		}

		userinfos = append(userinfos, retuserinfo)
		errinfo = append(errinfo, "")
	}
	// 返回结果
	if errflag {
		out.Errors = errinfo
		out.UserInfo = userinfos
		return out, errFailed
	}

	out.Success = true
	out.UserInfo = userinfos
	return out, nil
}
