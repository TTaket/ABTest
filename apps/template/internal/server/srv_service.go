package server

import (
	pb "ABTest/pkgs/proto/pb_experiment"
	"context"
)

// use service to implement the interface
func (s *xxxServer) GetExperiment(ctx context.Context, req *pb.GetExperimentRequest) (resp *pb.GetExperimentResponse, err error) {
	return s.xxxService.GetExperiment(ctx, req)
}

// use service to implement the interface
func (s *xxxServer) CreateExperiment(ctx context.Context, req *pb.CreateExperimentRequest) (resp *pb.CreateExperimentResponse, err error) {
	return s.xxxService.CreateExperiment(ctx, req)
}
