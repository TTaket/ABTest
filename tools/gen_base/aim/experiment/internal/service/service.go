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
	DeleteExperiment(ctx context.Context, in *pb.DeleteExperimentRequest, opts ...grpc.CallOption) (*pb.DeleteExperimentResponse, error)
	UpdateExperimentBaseInfo(ctx context.Context, in *pb.UpdateExperimentBaseInfoRequest, opts ...grpc.CallOption) (*pb.UpdateExperimentBaseInfoResponse, error)
	AddExperimentGroup(ctx context.Context, in *pb.AddExperimentGroupRequest, opts ...grpc.CallOption) (*pb.AddExperimentGroupResponse, error)
	DeleteExperimentGroup(ctx context.Context, in *pb.DeleteExperimentGroupRequest, opts ...grpc.CallOption) (*pb.DeleteExperimentGroupResponse, error)
	
}

type experimentService struct {

	// others..
}

func NewExperimentService() ExperimentService {
	return &experimentService{}
}
