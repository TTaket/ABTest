package service

import (
	"context"

	modelutils "ABTest/apps/Dao/userb/utils"
	"ABTest/pkgs/logger"
	pb "ABTest/pkgs/proto/pb_userb"
	"ABTest/pkgs/xmysql"
)

func (s *userbService) UpdateUser(ctx context.Context, in *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error) {
	logger.Infof("UpdateUser service begin: %v", in)

	out := &pb.UpdateUserResponse{
		Success: false,
	}
	db := xmysql.GetDB()
	if db == nil {
		out.Error = xmysql.ERR_GetDB_FAILED.Error()
		logger.Errorf("UpdateUser GetDB failed: %v", xmysql.ERR_GetDB_FAILED)
		return out, xmysql.ERR_GetDB_FAILED
	}
	usermodel, err := modelutils.TranslateProtoUserbInfoToUserbModel(in.UserInfo)
	if err != nil {
		out.Error = err.Error()
		logger.Errorf("UpdateUser translate failed: %v", err)
		return out, err
	}
	if err := usermodel.UpdateUserb(db); err != nil {
		out.Error = err.Error()
		logger.Errorf("UpdateUser failed: %v", err)
		return out, err
	}
	out.Success = true
	logger.Infof("UpdateUser service end: %v", out)
	return out, nil
}

func (s *userbService) BatchUpdateUser(ctx context.Context, in *pb.BatchUpdateUserRequest) (*pb.BatchUpdateUserResponse, error) {
	logger.Infof("BatchUpdateUser service begin: %v", in)

	out := &pb.BatchUpdateUserResponse{
		Success: false,
	}

	db := xmysql.GetDB()
	if db == nil {
		out.Errors = append(out.Errors, xmysql.ERR_GetDB_FAILED.Error())
		logger.Errorf("BatchUpdateUser GetDB failed: %v", xmysql.ERR_GetDB_FAILED)
		return out, xmysql.ERR_GetDB_FAILED
	}

	var errflag bool
	var errinfos []string

	for _, userinfo := range in.UserInfo {
		usermodel, err := modelutils.TranslateProtoUserbInfoToUserbModel(userinfo)
		if err != nil {
			errinfos = append(errinfos, err.Error())
			errflag = true
			continue
		}
		if err := usermodel.UpdateUserb(db); err != nil {
			errinfos = append(errinfos, err.Error())
			errflag = true
			continue
		}
		errinfos = append(errinfos, "")
	}
	if errflag {
		out.Errors = errinfos
		logger.Errorf("BatchUpdateUser service end with error: %v", out)
		return out, nil
	}
	out.Success = true
	logger.Infof("BatchUpdateUser service end: %v", out)
	return out, nil
}
