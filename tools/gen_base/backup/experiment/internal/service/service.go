package service

import (
	"context"

	pb "ABTest/pkgs/proto/pb_experiment"
)

// service 负责业务逻辑实现
type ExperimentService interface {
	GetExperiment(ctx context.Context, req *pb.GetExperimentRequest) (resp *pb.GetExperimentResponse, err error)
	CreateExperiment(ctx context.Context, req *pb.CreateExperimentRequest) (resp *pb.CreateExperimentResponse, err error)
}

type experimentService struct {

	// others..
}

func NewExperimentService() ExperimentService {
	return &experimentService{}
}
