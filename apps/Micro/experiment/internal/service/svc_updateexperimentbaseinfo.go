package service

import (
	"context"

	"ABTest/apps/Dao/experiment/utils"
	conf "ABTest/apps/Micro/experiment/internal/config"
	pb "ABTest/pkgs/proto/pb_experiment"
	mysql "ABTest/pkgs/xmysql"
)

// implement the interface
func (s *experimentService) UpdateExperimentBaseInfo(ctx context.Context, req *pb.UpdateExperimentBaseInfoRequest) (resp *pb.UpdateExperimentBaseInfoResponse, err error) {
	conf.Log.Infof("UpdateExperimentBaseInfo service begin: %v", req)

	// 这里实现逻辑
	// 1. 设置resp
	resp = &pb.UpdateExperimentBaseInfoResponse{
		Success: false,
	}
	// 2. db
	db := mysql.GetDB()
	if db == nil {
		conf.Log.Errorf("UpdateExperimentBaseInfo get db failed")
		resp.Error = "get db failed"
		return resp, mysql.ERR_GetDB_FAILED
	}

	// 3. 获取req 信息
	tx := db.Begin()
	experimentbaseinfo, experimentgroupsinfo := utils.TranslateProtoExperimentInfoToExperimentModel(req.ExperimentInfo)
	err = experimentbaseinfo.UpdateExperimentInfo(tx)
	if err != nil {
		conf.Log.Errorf("UpdateExperimentInfo failed: %v", err)
		tx.Rollback()
		resp.Error = err.Error()
		return resp, err
	}
	for i := range experimentgroupsinfo {
		groupPtr := &experimentgroupsinfo[i]
		err = groupPtr.UpdateExperimentGroup(tx)
		if err != nil {
			tx.Rollback()
			conf.Log.Errorf("UpdateExperimentGroup failed: %v", err)
			resp.Error = err.Error()
			return resp, err
		}
	}

	if err := tx.Commit().Error; err != nil {
		conf.Log.Errorf("Commit transaction failed: %v", err)
		resp.Error = "commit transaction failed"
		return resp, err
	}
	resp.Success = true
	conf.Log.Infof("UpdateExperimentBaseInfo service end: %v", resp)
	return resp, nil
}
