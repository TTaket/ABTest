package service

import (
	"context"

	utils "ABTest/apps/Dao/experiment/utils"
	conf "ABTest/apps/Micro/experiment/internal/config"
	pb "ABTest/pkgs/proto/pb_experiment"
	mysql "ABTest/pkgs/xmysql"
)

// implement the interface
func (s *experimentService) AddExperimentGroup(ctx context.Context, req *pb.AddExperimentGroupRequest) (resp *pb.AddExperimentGroupResponse, err error) {
	conf.Log.Infof("AddExperimentGroup service begin: %v", req)

	resp = &pb.AddExperimentGroupResponse{
		Success: false,
	}
	// 这里实现逻辑
	// 1. 连接数据库
	db := mysql.GetDB()
	if db == nil {
		conf.Log.Errorf("AddExperimentGroup get db failed")
		resp.Error = "get db failed"
		return resp, err
	}

	// 2. 解析请求参数
	groupinfo := utils.TranslateProtoExperimentGroupToExperimentModel(req.Group)

	// 3. 创建实验组
	err = groupinfo.CreateExperimentGroup(db)
	if err != nil {
		conf.Log.Errorf("CreateExperimentGroup failed: %v", err)
		resp.Error = err.Error()
		return resp, err
	}

	// 4. 返回结果
	resp.Success = true
	conf.Log.Infof("AddExperimentGroup service end: %v", resp)
	return resp, nil
}
