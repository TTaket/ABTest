package service

import (
	"context"
	"errors"

	model "ABTest/apps/Dao/userb/model"
	pb "ABTest/pkgs/proto/pb_userb"
	xmysql "ABTest/pkgs/xmysql"
)

func (s *userbService) DeleteUser(ctx context.Context, in *pb.DeleteUserRequest) (out *pb.DeleteUserResponse, reterr error) {
	// 创建返回体
	out = &pb.DeleteUserResponse{
		Success: false,
	}

	// 取出id
	id := in.UserId

	// 获得db
	db := xmysql.GetDB()
	if db == nil {
		return nil, xmysql.ERR_GetDB_FAILED
	}

	// 删除用户
	err := model.DeleteUserbByID(db, id)
	// 返回结果
	if err != nil {
		out.Error = err.Error()
		return out, err
	}

	out.Success = true
	return out, nil
}

func (s *userbService) BatchDeleteUser(ctx context.Context, in *pb.BatchDeleteUserRequest) (out *pb.BatchDeleteUserResponse, reterr error) {
	// 创建返回体
	out = &pb.BatchDeleteUserResponse{
		Success: false,
	}

	// 取出ids
	ids := in.UserId

	// 获得db
	db := xmysql.GetDB()
	if db == nil {
		out.Errors = append(out.Errors, xmysql.ERR_GetDB_FAILED.Error())
		return nil, xmysql.ERR_GetDB_FAILED
	}

	// 设置错误标志
	var errflag bool = false
	var errinfo []string = []string{}

	// 删除用户
	for _, id := range ids {
		err := model.DeleteUserbByID(db, id)
		if err != nil {
			errflag = true
			errinfo = append(errinfo, err.Error())
		}
		errinfo = append(errinfo, "")
	}

	// 返回结果
	if errflag {
		out.Success = false
		out.Errors = errinfo
		return out, errors.New("BatchDeleteUser failed")
	}
	out.Success = true
	return out, nil
}
