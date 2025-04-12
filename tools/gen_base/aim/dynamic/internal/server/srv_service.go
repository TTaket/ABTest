package server

import (
	pb "ABTest/pkgs/proto/pb_dynamic"
	"context"
)

// service 负责业务逻辑实现

func (s *dynamicServer) CreateLayerUserBinding(ctx context.Context, req *pb.CreateLayerUserBindingRequest) (resp *pb.CreateLayerUserBindingResponse, err error) {
	return s.dynamicService.CreateLayerUserBinding(ctx, req)
}

	
func (s *dynamicServer) GetLayerUserBinding(ctx context.Context, req *pb.GetLayerUserBindingRequest) (resp *pb.GetLayerUserBindingResponse, err error) {
	return s.dynamicService.GetLayerUserBinding(ctx, req)
}

	
func (s *dynamicServer) DeleteLayerUserBinding(ctx context.Context, req *pb.DeleteLayerUserBindingRequest) (resp *pb.DeleteLayerUserBindingResponse, err error) {
	return s.dynamicService.DeleteLayerUserBinding(ctx, req)
}

	
func (s *dynamicServer) CreateExperimentUserBinding(ctx context.Context, req *pb.CreateExperimentUserBindingRequest) (resp *pb.CreateExperimentUserBindingResponse, err error) {
	return s.dynamicService.CreateExperimentUserBinding(ctx, req)
}

	
func (s *dynamicServer) GetExperimentUserBinding(ctx context.Context, req *pb.GetExperimentUserBindingRequest) (resp *pb.GetExperimentUserBindingResponse, err error) {
	return s.dynamicService.GetExperimentUserBinding(ctx, req)
}

	
func (s *dynamicServer) DeleteExperimentUserBinding(ctx context.Context, req *pb.DeleteExperimentUserBindingRequest) (resp *pb.DeleteExperimentUserBindingResponse, err error) {
	return s.dynamicService.DeleteExperimentUserBinding(ctx, req)
}

	
func (s *dynamicServer) CreateExperimentLayerBinding(ctx context.Context, req *pb.CreateExperimentLayerBindingRequest) (resp *pb.CreateExperimentLayerBindingResponse, err error) {
	return s.dynamicService.CreateExperimentLayerBinding(ctx, req)
}

	
func (s *dynamicServer) GetExperimentLayerBinding(ctx context.Context, req *pb.GetExperimentLayerBindingRequest) (resp *pb.GetExperimentLayerBindingResponse, err error) {
	return s.dynamicService.GetExperimentLayerBinding(ctx, req)
}

	
func (s *dynamicServer) GetLayerExperiments(ctx context.Context, req *pb.GetLayerExperimentsRequest) (resp *pb.GetLayerExperimentsResponse, err error) {
	return s.dynamicService.GetLayerExperiments(ctx, req)
}

	
func (s *dynamicServer) DeleteExperimentLayerBinding(ctx context.Context, req *pb.DeleteExperimentLayerBindingRequest) (resp *pb.DeleteExperimentLayerBindingResponse, err error) {
	return s.dynamicService.DeleteExperimentLayerBinding(ctx, req)
}

	
