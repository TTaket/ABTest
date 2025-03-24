package service

import (
	"context"

	"ABTest/apps/Dao/experiment/model"
	conf "ABTest/apps/Micro/experiment/internal/config"
	pb "ABTest/pkgs/proto/pb_experiment"
	mysql "ABTest/pkgs/xmysql"
)

// implement the interface
func (s *experimentService) DeleteExperimentGroup(ctx context.Context, req *pb.DeleteExperimentGroupRequest) (resp *pb.DeleteExperimentGroupResponse, err error) {
	conf.Log.Infof("DeleteExperimentGroup service begin: %v", req)

	// 这里实现逻辑
	// 1. 设置resp
	resp = &pb.DeleteExperimentGroupResponse{
		Success: false,
	}
	// 2. 连接数据库
	db := mysql.GetDB()
	if db == nil {
		conf.Log.Errorf("DeleteExperimentGroup get db failed")
		resp.Error = "get db failed"
		return resp, mysql.ERR_GetDB_FAILED
	}
	// 3. 删除实验组
	err = model.DeleteExperimentGroupByID(db, req.GroupId)
	if err != nil {
		conf.Log.Errorf("DeleteExperimentGroupByID failed: %v", err)
		resp.Error = err.Error()
		return resp, err
	}
	resp.Success = true
	conf.Log.Infof("DeleteExperimentGroup service end: %v", resp)
	return resp, nil
}
