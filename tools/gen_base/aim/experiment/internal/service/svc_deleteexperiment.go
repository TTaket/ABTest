package service

import (
	"context"

	conf "ABTest/apps/Micro/experiment/internal/config"
	pb "ABTest/pkgs/proto/pb_experiment"
)

// implement the interface
func (s *experimentService) DeleteExperiment(ctx context.Context, req *pb.DeleteExperimentRequest) (resp *pb.DeleteExperimentResponse, err error) {
	conf.Log.Infof("DeleteExperiment service begin: %v", req)

	// 这里实现逻辑

	
	conf.Log.Infof("DeleteExperiment service end: %v", resp)
	return resp, nil
}
