package service

import (
	"context"
	"errors"

	modelutils "ABTest/apps/Dao/userb/utils"
	pb "ABTest/pkgs/proto/pb_userb"
	xmysql "ABTest/pkgs/xmysql"
)

func (s *userbService) AddUser(ctx context.Context, in *pb.AddUserRequest) (*pb.AddUserResponse, error) {
	// 创建返回体
	out := &pb.AddUserResponse{
		Success: false,
		UserId:  0,
	}

	// 获得用户信息
	userbmodel, err := modelutils.TranslateProtoUserbInfoToUserbModel(in.UserInfo)
	if err != nil {
		out.Error = err.Error()
		return out, err
	}

	// get db
	db := xmysql.GetDB()
	if db == nil {
		out.Error = xmysql.ERR_GetDB_FAILED.Error()
		return out, xmysql.ERR_GetDB_FAILED
	}

	// 创建用户
	err = userbmodel.CreateUserb(db)
	if err != nil {
		out.Error = err.Error()
		return out, err
	}
	//Todo: temp
	// 调用ExperimentClient的AddExperiment方法
	/* 	{
		experimentClient := client.NewExperimentClient()
		// get experiment
		{
			req := &pb_experiment.GetExperimentRequest{
				ExperimentId: "1",
			}
			resp, err := experimentClient.GetExperiment(ctx, req)
			if err != nil {
				panic(err)
			}
			fmt.Printf("TestDemo: get experiment %s\n", resp.String())
		}
	} */

	out.Success = true
	out.UserId = userbmodel.UserID
	return out, nil
}

func (s *userbService) BatchAddUser(ctx context.Context, in *pb.BatchAddUserRequest) (*pb.BatchAddUserResponse, error) {
	// 创建返回体
	out := &pb.BatchAddUserResponse{
		Success: false,
	}

	var userIDs []uint64
	var errinfos []string
	var errflag bool

	// get db
	db := xmysql.GetDB()
	if db == nil {
		out.Errors = append(out.Errors, xmysql.ERR_GetDB_FAILED.Error())
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

	// 返回结果
	if errflag {
		err := errors.New("BatchAddUser failed")
		out.Success = false
		out.Errors = errinfos
		out.UserId = userIDs
		return out, err
	}
	out.Success = true
	out.UserId = userIDs
	return out, nil
}
