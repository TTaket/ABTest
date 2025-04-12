package server

import (
	pb "ABTest/pkgs/proto/pb_layer"
	"context"
)

// service 负责业务逻辑实现

func (s *layerServer) CreateLayer(ctx context.Context, req *pb.CreateLayerRequest) (resp *pb.CreateLayerResponse, err error) {
	return s.layerService.CreateLayer(ctx, req)
}

	
func (s *layerServer) GetLayer(ctx context.Context, req *pb.GetLayerRequest) (resp *pb.GetLayerResponse, err error) {
	return s.layerService.GetLayer(ctx, req)
}

	
func (s *layerServer) UpdateLayer(ctx context.Context, req *pb.UpdateLayerRequest) (resp *pb.UpdateLayerResponse, err error) {
	return s.layerService.UpdateLayer(ctx, req)
}

	
func (s *layerServer) DeleteLayer(ctx context.Context, req *pb.DeleteLayerRequest) (resp *pb.DeleteLayerResponse, err error) {
	return s.layerService.DeleteLayer(ctx, req)
}

	
func (s *layerServer) ListLayers(ctx context.Context, req *pb.ListLayersRequest) (resp *pb.ListLayersResponse, err error) {
	return s.layerService.ListLayers(ctx, req)
}

	
