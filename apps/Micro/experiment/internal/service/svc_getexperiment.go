package service

import (
	"context"

	model "ABTest/apps/Dao/experiment/model"
	"ABTest/apps/Dao/experiment/utils"
	conf "ABTest/apps/Micro/experiment/internal/config"
	pb "ABTest/pkgs/proto/pb_experiment"
	mysql "ABTest/pkgs/xmysql"
)

// implement the interface
func (s *experimentService) GetExperiment(ctx context.Context, req *pb.GetExperimentRequest) (resp *pb.GetExperimentResponse, err error) {
	conf.Log.Infof("GetExperiment service begin: %v", req)

	// 这里实现逻辑
	// 1. 设置resp
	resp = &pb.GetExperimentResponse{
		Success: false,
	}
	// 2. db
	db := mysql.GetDB()
	if db == nil {
		conf.Log.Errorf("GetExperiment get db failed")
		resp.Error = "get db failed"
		return resp, mysql.ERR_GetDB_FAILED
	}
	// 3. 查询实验
	// 3.1 查询实验
	experimentinfobasic, err := model.GetExperimentInfoByID(db, req.ExperimentId)
	if err != nil {
		conf.Log.Errorf("GetExperimentInfoByID failed: %v", err)
		resp.Error = err.Error()
		return resp, err
	}
	// 3.2 查询实验组
	experimentgroupinfo, err := model.GetExperimentGroupsByExperimentID(db, req.ExperimentId)
	if err != nil {
		conf.Log.Errorf("GetExperimentGroupsByExperimentID failed: %v", err)
		resp.Error = err.Error()
		return resp, err
	}
	// 组装
	experimentinfo := utils.TranslateExperimentModelToProtoExperimentInfo(*experimentinfobasic, experimentgroupinfo)

	// 返回
	resp.Success = true
	resp.ExperimentInfo = experimentinfo

	conf.Log.Infof("GetExperiment service end: %v", resp)
	return resp, nil
}
