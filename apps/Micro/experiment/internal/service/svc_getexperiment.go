package service

import (
	"context"

	conf "ABTest/apps/Micro/experiment/internal/config"
	pb "ABTest/pkgs/proto/pb_experiment"
)

// implement the interface
func (s *experimentService) GetExperiment(ctx context.Context, req *pb.GetExperimentRequest) (resp *pb.GetExperimentResponse, err error) {
	conf.Log.Infof("GetExperiment service begin: %v", req)

	// 这里实现逻辑

	conf.Log.Infof("GetExperiment service end: %v", resp)
	return resp, nil
}
