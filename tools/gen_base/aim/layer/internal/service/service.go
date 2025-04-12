package service

import (
	"context"

	pb "ABTest/pkgs/proto/pb_layer"
		
	"google.golang.org/grpc"
)

// service 负责业务逻辑实现
type LayerService interface {
	CreateLayer(ctx context.Context, in *pb.CreateLayerRequest) (*pb.CreateLayerResponse, error)
	GetLayer(ctx context.Context, in *pb.GetLayerRequest) (*pb.GetLayerResponse, error)
	UpdateLayer(ctx context.Context, in *pb.UpdateLayerRequest) (*pb.UpdateLayerResponse, error)
	DeleteLayer(ctx context.Context, in *pb.DeleteLayerRequest) (*pb.DeleteLayerResponse, error)
	ListLayers(ctx context.Context, in *pb.ListLayersRequest) (*pb.ListLayersResponse, error)
	
}

type layerService struct {

	// others..
}

func NewLayerService() LayerService {
	return &layerService{}
}
