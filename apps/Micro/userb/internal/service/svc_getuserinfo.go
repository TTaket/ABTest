package service

import (
	"context"
	"errors"

	model "ABTest/apps/Dao/userb/model"
	modelutils "ABTest/apps/Dao/userb/utils"
	"ABTest/pkgs/logger"
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
	logger.Infof("GetUserInfo service begin: %v", in)

	out := &pb.GetUserInfoResponse{
		Success: false,
	}
	db := xmysql.GetDB()
	if db == nil {
		out.Error = errGetDBFailed.Error()
		logger.Errorf("GetUserInfo GetDB failed: %v", errGetDBFailed)
		return out, errGetDBFailed
	}

	userbmodel, err := model.GetUserbByID(db, *in.UserInfo.UserId)
	if err != nil {
		out.Error = err.Error()
		logger.Errorf("GetUserInfo failed: %v", err)
		return out, err
	}

	userinfo, err := modelutils.TranslateUserbModelToProtoUserbInfo(userbmodel)
	if err != nil {
		out.Error = err.Error()
		logger.Errorf("GetUserInfo translate failed: %v", err)
		return out, err
	}

	out.Success = true
	out.UserInfo = userinfo
	logger.Infof("GetUserInfo service end: %v", out)
	return out, nil
}

func (s *userbService) BatchGetUserInfo(ctx context.Context, in *pb.BatchGetUserInfoRequest) (*pb.BatchGetUserInfoResponse, error) {
	logger.Infof("BatchGetUserInfo service begin: %v", in)

	out := &pb.BatchGetUserInfoResponse{
		Success: false,
	}
	var errflag bool = false
	var userinfos []*pb.UserInfo
	var errinfo []string

	db := xmysql.GetDB()
	if db == nil {
		out.Errors = append(out.Errors, errGetDBFailed.Error())
		logger.Errorf("BatchGetUserInfo GetDB failed: %v", errGetDBFailed)
		return out, errGetDBFailed
	}

	if in.UserInfo == nil {
		out.Errors = append(out.Errors, errIdRequired.Error())
		logger.Errorf("BatchGetUserInfo missing user id: %v", errIdRequired)
		return out, errIdRequired
	}

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

	if errflag {
		out.Errors = errinfo
		out.UserInfo = userinfos
		logger.Errorf("BatchGetUserInfo service end with error: %v", out)
		return out, errFailed
	}

	out.Success = true
	out.UserInfo = userinfos
	logger.Infof("BatchGetUserInfo service end: %v", out)
	return out, nil
}
