package service

import (
	"context"

	conf "ABTest/apps/Micro/experiment/internal/config"
	pb "ABTest/pkgs/proto/pb_experiment"
)

// implement the interface
func (s *experimentService) CreateExperiment(ctx context.Context, req *pb.CreateExperimentRequest) (resp *pb.CreateExperimentResponse, err error) {
	conf.Log.Infof("CreateExperiment service begin: %v", req)

	// 这里实现逻辑

	
	conf.Log.Infof("CreateExperiment service end: %v", resp)
	return resp, nil
}
