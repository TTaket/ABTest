package service

import (
	"context"

	model "ABTest/apps/Dao/experiment/model"
	conf "ABTest/apps/Micro/experiment/internal/config"
	pb "ABTest/pkgs/proto/pb_experiment"
	mysql "ABTest/pkgs/xmysql"
)

// implement the interface
func (s *experimentService) DeleteExperiment(ctx context.Context, req *pb.DeleteExperimentRequest) (resp *pb.DeleteExperimentResponse, err error) {
	conf.Log.Infof("DeleteExperiment service begin: %v", req)

	// 这里实现逻辑
	// 1. 设置resp
	resp = &pb.DeleteExperimentResponse{
		Success: false,
	}
	// 2. 连接数据库
	db := mysql.GetDB()
	if db == nil {
		conf.Log.Errorf("DeleteExperiment get db failed")
		resp.Error = "get db failed"
		return resp, mysql.ERR_GetDB_FAILED
	}
	// 3. 删除实验
	// 开启事务
	tx := db.Begin()
	if tx.Error != nil {
		conf.Log.Errorf("Begin transaction failed: %v", tx.Error)
		resp.Error = "begin transaction failed"
		return resp, tx.Error
	}

	// 删除实验组
	err = model.DeleteExperimentInfoByID(tx, req.ExperimentId)
	if err != nil {
		conf.Log.Errorf("DeleteExperimentInfoByID failed: %v", err)
		tx.Rollback()
		resp.Error = err.Error()
		return resp, err
	}

	// 删除实验组下的实验
	err = model.DeleteExperimentGroupsByExperimentID(tx, req.ExperimentId)
	if err != nil {
		conf.Log.Errorf("DeleteExperimentGroupsByExperimentID failed: %v", err)
		tx.Rollback()
		resp.Error = err.Error()
		return resp, err
	}

	// 提交事务
	if err = tx.Commit().Error; err != nil {
		conf.Log.Errorf("Commit transaction failed: %v", err)
		resp.Error = "commit transaction failed"
		return resp, err
	}

	// 5. 返回结果
	resp.Success = true
	conf.Log.Infof("DeleteExperiment service end: %v", resp)
	return resp, nil
}
