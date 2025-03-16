package server

import (
	pb "ABTest/pkgs/proto/pb_experiment"
	"context"
)

// service 负责业务逻辑实现

func (s *experimentServer) CreateExperiment(ctx context.Context, req *pb.CreateExperimentRequest) (resp *pb.CreateExperimentResponse, err error) {
	return s.experimentService.CreateExperiment(ctx, req)
}

	
func (s *experimentServer) GetExperiment(ctx context.Context, req *pb.GetExperimentRequest) (resp *pb.GetExperimentResponse, err error) {
	return s.experimentService.GetExperiment(ctx, req)
}

	
