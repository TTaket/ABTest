package service

import (
	"context"

	pb "ABTest/pkgs/proto/pb_dynamic"
		
	"google.golang.org/grpc"
)

// service 负责业务逻辑实现
type DynamicService interface {
	CreateLayerUserBinding(ctx context.Context, in *pb.CreateLayerUserBindingRequest) (*pb.CreateLayerUserBindingResponse, error)
	GetLayerUserBinding(ctx context.Context, in *pb.GetLayerUserBindingRequest) (*pb.GetLayerUserBindingResponse, error)
	DeleteLayerUserBinding(ctx context.Context, in *pb.DeleteLayerUserBindingRequest) (*pb.DeleteLayerUserBindingResponse, error)
	CreateExperimentUserBinding(ctx context.Context, in *pb.CreateExperimentUserBindingRequest) (*pb.CreateExperimentUserBindingResponse, error)
	GetExperimentUserBinding(ctx context.Context, in *pb.GetExperimentUserBindingRequest) (*pb.GetExperimentUserBindingResponse, error)
	DeleteExperimentUserBinding(ctx context.Context, in *pb.DeleteExperimentUserBindingRequest) (*pb.DeleteExperimentUserBindingResponse, error)
	CreateExperimentLayerBinding(ctx context.Context, in *pb.CreateExperimentLayerBindingRequest) (*pb.CreateExperimentLayerBindingResponse, error)
	GetExperimentLayerBinding(ctx context.Context, in *pb.GetExperimentLayerBindingRequest) (*pb.GetExperimentLayerBindingResponse, error)
	GetLayerExperiments(ctx context.Context, in *pb.GetLayerExperimentsRequest) (*pb.GetLayerExperimentsResponse, error)
	DeleteExperimentLayerBinding(ctx context.Context, in *pb.DeleteExperimentLayerBindingRequest) (*pb.DeleteExperimentLayerBindingResponse, error)
	
}

type dynamicService struct {

	// others..
}

func NewDynamicService() DynamicService {
	return &dynamicService{}
}
