package service

import (
	"context"

	pb "ABTest/pkgs/proto/pb_experiment"
)

// service 负责业务逻辑实现
type ExperimentService interface {
	CreateExperiment(ctx context.Context, in *pb.CreateExperimentRequest ) (*pb.CreateExperimentResponse, error)
	GetExperiment(ctx context.Context, in *pb.GetExperimentRequest ) (*pb.GetExperimentResponse, error)
	DeleteExperiment(ctx context.Context, in *pb.DeleteExperimentRequest ) (*pb.DeleteExperimentResponse, error)
	UpdateExperimentBaseInfo(ctx context.Context, in *pb.UpdateExperimentBaseInfoRequest ) (*pb.UpdateExperimentBaseInfoResponse, error)
	AddExperimentGroup(ctx context.Context, in *pb.AddExperimentGroupRequest ) (*pb.AddExperimentGroupResponse, error)
	DeleteExperimentGroup(ctx context.Context, in *pb.DeleteExperimentGroupRequest ) (*pb.DeleteExperimentGroupResponse, error)
	
}

type experimentService struct {

	// others..
}

func NewExperimentService() ExperimentService {
	return &experimentService{}
}
