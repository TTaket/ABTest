package service

import (
	"context"

	modelutil "ABTest/apps/Dao/userb/utils"
	pb "ABTest/pkgs/proto/pb_userb"
	"ABTest/pkgs/xmysql"
)

func (s *userbService) UpdateUser(ctx context.Context, in *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error) {
	out := &pb.UpdateUserResponse{
		Success: false,
	}
	db := xmysql.GetDB()
	if db == nil {
		out.Error = xmysql.ERR_GetDB_FAILED.Error()
		return out, xmysql.ERR_GetDB_FAILED
	}
	usermodel, err := modelutil.TranslateProtoUserbInfoToUserbModel(in.UserInfo)
	if err != nil {
		out.Error = err.Error()
		return out, err
	}
	if err := usermodel.UpdateUserb(db); err != nil {
		out.Error = err.Error()
		return out, err
	}
	out.Success = true
	return out, nil
}
func (s *userbService) BatchUpdateUser(ctx context.Context, in *pb.BatchUpdateUserRequest) (*pb.BatchUpdateUserResponse, error) {

	out := &pb.BatchUpdateUserResponse{
		Success: false,
	}

	db := xmysql.GetDB()
	if db == nil {
		out.Errors = append(out.Errors, xmysql.ERR_GetDB_FAILED.Error())
		return out, xmysql.ERR_GetDB_FAILED
	}

	var errflag bool
	var errinfos []string

	for _, userinfo := range in.UserInfo {
		usermodel, err := modelutil.TranslateProtoUserbInfoToUserbModel(userinfo)
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
		return out, nil
	}
	out.Success = true
	return out, nil
}
