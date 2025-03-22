package server

import (
	pb "ABTest/pkgs/proto/pb_experiment"
	"context"
)

// use service to implement the interface
func (s *experimentServer) GetExperiment(ctx context.Context, req *pb.GetExperimentRequest) (resp *pb.GetExperimentResponse, err error) {
	return s.experimentService.GetExperiment(ctx, req)
}

// use service to implement the interface
func (s *experimentServer) CreateExperiment(ctx context.Context, req *pb.CreateExperimentRequest) (resp *pb.CreateExperimentResponse, err error) {
	return s.experimentService.CreateExperiment(ctx, req)
}
