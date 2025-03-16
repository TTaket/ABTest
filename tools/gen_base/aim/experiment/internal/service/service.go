package service

import (
	"context"

	pb "ABTest/pkgs/proto/pb_experiment"
		
	"google.golang.org/grpc"
)

// service 负责业务逻辑实现
type ExperimentService interface {
	CreateExperiment(ctx context.Context, in *pb.CreateExperimentRequest, opts ...grpc.CallOption) (*pb.CreateExperimentResponse, error)
	GetExperiment(ctx context.Context, in *pb.GetExperimentRequest, opts ...grpc.CallOption) (*pb.GetExperimentResponse, error)
	
}

type experimentService struct {

	// others..
}

func NewExperimentService() ExperimentService {
	return &experimentService{}
}
