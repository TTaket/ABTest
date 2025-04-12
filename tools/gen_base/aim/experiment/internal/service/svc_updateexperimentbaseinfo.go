package service

import (
	"context"

	conf "ABTest/apps/Micro/experiment/internal/config"
	pb "ABTest/pkgs/proto/pb_experiment"
)

// implement the interface
func (s *experimentService) UpdateExperimentBaseInfo(ctx context.Context, req *pb.UpdateExperimentBaseInfoRequest) (resp *pb.UpdateExperimentBaseInfoResponse, err error) {
	conf.Log.Infof("UpdateExperimentBaseInfo service begin: %v", req)

	// 这里实现逻辑

	
	conf.Log.Infof("UpdateExperimentBaseInfo service end: %v", resp)
	return resp, nil
}
