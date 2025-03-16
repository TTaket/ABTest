package service

import (
	"context"
	"errors"

	model "ABTest/apps/Dao/userb/model"
	"ABTest/pkgs/logger"
	pb "ABTest/pkgs/proto/pb_userb"
	xmysql "ABTest/pkgs/xmysql"
)

func (s *userbService) DeleteUser(ctx context.Context, in *pb.DeleteUserRequest) (out *pb.DeleteUserResponse, reterr error) {
	logger.Infof("DeleteUser service begin: %v", in)

	out = &pb.DeleteUserResponse{
		Success: false,
	}

	id := in.UserId
	db := xmysql.GetDB()
	if db == nil {
		logger.Errorf("DeleteUser GetDB failed: %v", xmysql.ERR_GetDB_FAILED)
		return nil, xmysql.ERR_GetDB_FAILED
	}

	err := model.DeleteUserbByID(db, id)
	if err != nil {
		out.Error = err.Error()
		logger.Errorf("DeleteUser failed: %v", err)
		return out, err
	}

	out.Success = true
	logger.Infof("DeleteUser service end: %v", out)
	return out, nil
}

func (s *userbService) BatchDeleteUser(ctx context.Context, in *pb.BatchDeleteUserRequest) (out *pb.BatchDeleteUserResponse, reterr error) {
	logger.Infof("BatchDeleteUser service begin: %v", in)

	out = &pb.BatchDeleteUserResponse{
		Success: false,
	}
	ids := in.UserId
	db := xmysql.GetDB()
	if db == nil {
		out.Errors = append(out.Errors, xmysql.ERR_GetDB_FAILED.Error())
		logger.Errorf("BatchDeleteUser GetDB failed: %v", xmysql.ERR_GetDB_FAILED)
		return nil, xmysql.ERR_GetDB_FAILED
	}

	var errflag bool = false
	var errinfo []string

	for _, id := range ids {
		err := model.DeleteUserbByID(db, id)
		if err != nil {
			errflag = true
			errinfo = append(errinfo, err.Error())
		} else {
			errinfo = append(errinfo, "")
		}
	}

	if errflag {
		out.Success = false
		out.Errors = errinfo
		logger.Errorf("BatchDeleteUser service end with error: %v", out)
		return out, errors.New("BatchDeleteUser failed")
	}
	out.Success = true
	logger.Infof("BatchDeleteUser service end: %v", out)
	return out, nil
}
