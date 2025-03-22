package service

import (
	"context"
	"errors"

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
	// 如果试图修改用户的ID，返回错误
	if in == nil {
		out.Error = errors.New("input is nil").Error()
		logger.Errorf("UpdateUser failed: %v", errors.New("input is nil"))
		return out, errors.New("input is nil")
	}
	if *in.UserInfo.UserId != in.UserId {
		out.Error = errors.New("cannot update user id").Error()
		logger.Errorf("UpdateUser failed: %v", errors.New("cannot update user id"))
		return out, errors.New("cannot update user id")
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

	if in == nil || (len(in.UserInfo) != len(in.UserId)) {
		out.Errors = append(out.Errors, errors.New("input is error ").Error())
		logger.Errorf("BatchUpdateUser failed: %v", errors.New("input is error "))
		return out, errors.New("input is error")
	}

	var errflag bool
	var errinfos []string

	for i, userinfo := range in.UserInfo {
		// 如果试图修改用户的ID，返回错误
		if userinfo == nil {
			errinfos = append(errinfos, errors.New("input is nil").Error())
			errflag = true
			continue
		}
		if *userinfo.UserId != in.UserId[i] {
			errinfos = append(errinfos, errors.New("cannot update user id").Error())
			errflag = true
			continue
		}

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
