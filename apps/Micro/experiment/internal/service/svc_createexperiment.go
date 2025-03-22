package service

import (
	"context"

	"ABTest/apps/Dao/experiment/utils"
	conf "ABTest/apps/Micro/experiment/internal/config"
	pb "ABTest/pkgs/proto/pb_experiment"
	mysql "ABTest/pkgs/xmysql"
)

// implement the interface
func (s *experimentService) CreateExperiment(ctx context.Context, req *pb.CreateExperimentRequest) (resp *pb.CreateExperimentResponse, err error) {
	conf.Log.Infof("CreateExperiment service begin: %v", req)

	// 这里实现逻辑
	resp = &pb.CreateExperimentResponse{
		Success: false,
	}
	// 1. 连接数据库
	db := mysql.GetDB()
	if db == nil {
		conf.Log.Errorf("CreateExperiment get db failed")
		resp.Error = "get db failed"
		return nil, err
	}

	// 2. 解析请求参数
	// model.ExperimentInfoBasic, []model.ExperimentGroupBasic
	infobasic, groupsbasic := utils.TranslateProtoExperimentInfoToExperimentModel(req.ExperimentInfo)

	// 3. 创建实验
	err = infobasic.CreateExperimentInfo(db)
	if err != nil {
		conf.Log.Errorf("CreateExperimentInfo failed: %v", err)
		resp.Error = err.Error()
		return nil, err
	}

	create_group_err := false
	create_group_err_msg := []string{}
	// 4. 创建实验组
	for _, group := range groupsbasic {
		err = group.CreateExperimentGroup(db)
		if err != nil {
			create_group_err = true
			create_group_err_msg = append(create_group_err_msg, err.Error())
		} else {
			create_group_err_msg = append(create_group_err_msg, "")
		}
	}
	if create_group_err {
		// 日志
		conf.Log.Errorf("CreateExperimentGroup failed ")
		for i, msg := range create_group_err_msg {
			if msg == "" {
				continue
			}
			conf.Log.Errorf("CreateExperimentGroup failed: groupinfo : %v , error : %v", groupsbasic[i], msg)
		}
		resp.Error = err.Error()
		return resp, err
	}

	resp = &pb.CreateExperimentResponse{
		Success:        true,
		ExperimentId:   infobasic.ExperimentID,
		ExperimentInfo: utils.TranslateExperimentModelToProtoExperimentInfo(infobasic, groupsbasic),
	}

	conf.Log.Infof("CreateExperiment service end: %v", resp)
	return resp, nil
}
