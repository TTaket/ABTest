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

	
func (s *experimentServer) DeleteExperiment(ctx context.Context, req *pb.DeleteExperimentRequest) (resp *pb.DeleteExperimentResponse, err error) {
	return s.experimentService.DeleteExperiment(ctx, req)
}

	
func (s *experimentServer) UpdateExperimentBaseInfo(ctx context.Context, req *pb.UpdateExperimentBaseInfoRequest) (resp *pb.UpdateExperimentBaseInfoResponse, err error) {
	return s.experimentService.UpdateExperimentBaseInfo(ctx, req)
}

	
func (s *experimentServer) AddExperimentGroup(ctx context.Context, req *pb.AddExperimentGroupRequest) (resp *pb.AddExperimentGroupResponse, err error) {
	return s.experimentService.AddExperimentGroup(ctx, req)
}

	
func (s *experimentServer) DeleteExperimentGroup(ctx context.Context, req *pb.DeleteExperimentGroupRequest) (resp *pb.DeleteExperimentGroupResponse, err error) {
	return s.experimentService.DeleteExperimentGroup(ctx, req)
}

	
