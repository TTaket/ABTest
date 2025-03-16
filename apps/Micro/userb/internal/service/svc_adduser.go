package service

import (
	"context"
	"errors"

	modelutils "ABTest/apps/Dao/userb/utils"
	"ABTest/pkgs/logger"
	pb "ABTest/pkgs/proto/pb_userb"
	xmysql "ABTest/pkgs/xmysql"
)

func (s *userbService) AddUser(ctx context.Context, in *pb.AddUserRequest) (*pb.AddUserResponse, error) {
	logger.Infof("AddUser service begin: %v", in)

	out := &pb.AddUserResponse{
		Success: false,
		UserId:  0,
	}

	userbmodel, err := modelutils.TranslateProtoUserbInfoToUserbModel(in.UserInfo)
	if err != nil {
		logger.Errorf("TranslateProtoUserbInfoToUserbModel failed: %v", err)
		out.Error = err.Error()
		return out, err
	}

	db := xmysql.GetDB()
	if db == nil {
		logger.Errorf("GetDB failed")
		out.Error = xmysql.ERR_GetDB_FAILED.Error()
		return out, xmysql.ERR_GetDB_FAILED
	}

	err = userbmodel.CreateUserb(db)
	if err != nil {
		logger.Errorf("CreateUserb failed: %v", err)
		out.Error = err.Error()
		return out, err
	}

	out.Success = true
	out.UserId = userbmodel.UserID
	logger.Infof("AddUser service end: %v", out)
	return out, nil
}

func (s *userbService) BatchAddUser(ctx context.Context, in *pb.BatchAddUserRequest) (*pb.BatchAddUserResponse, error) {
	logger.Infof("BatchAddUser service begin: %v", in)

	out := &pb.BatchAddUserResponse{
		Success: false,
	}
	var userIDs []uint64
	var errinfos []string
	var errflag bool

	db := xmysql.GetDB()
	if db == nil {
		out.Errors = append(out.Errors, xmysql.ERR_GetDB_FAILED.Error())
		logger.Errorf("BatchAddUser GetDB failed: %v", xmysql.ERR_GetDB_FAILED)
		return out, xmysql.ERR_GetDB_FAILED
	}

	for _, userInfo := range in.UserInfo {
		userbmodel, err := modelutils.TranslateProtoUserbInfoToUserbModel(userInfo)
		if err != nil {
			errflag = true
			userIDs = append(userIDs, 0)
			errinfos = append(errinfos, err.Error())
			continue
		}

		err = userbmodel.CreateUserb(db)
		if err != nil {
			errflag = true
			userIDs = append(userIDs, 0)
			errinfos = append(errinfos, err.Error())
			continue
		}

		userIDs = append(userIDs, userbmodel.UserID)
		errinfos = append(errinfos, "")
	}

	if errflag {
		err := errors.New("BatchAddUser failed")
		out.Success = false
		out.Errors = errinfos
		out.UserId = userIDs
		logger.Errorf("BatchAddUser service end with error: %v", out)
		return out, err
	}
	out.Success = true
	out.UserId = userIDs
	logger.Infof("BatchAddUser service end: %v", out)
	return out, nil
}
