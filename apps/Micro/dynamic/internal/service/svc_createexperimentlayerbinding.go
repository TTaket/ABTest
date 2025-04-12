package service

import (
	"context"

	conf "ABTest/apps/Micro/dynamic/internal/config"
	pb "ABTest/pkgs/proto/pb_dynamic"
)

// implement the interface
func (s *dynamicService) CreateExperimentLayerBinding(ctx context.Context, req *pb.CreateExperimentLayerBindingRequest) (resp *pb.CreateExperimentLayerBindingResponse, err error) {
	conf.Log.Infof("CreateExperimentLayerBinding service begin: %v", req)

	// 这里实现逻辑

	conf.Log.Infof("CreateExperimentLayerBinding service end: %v", resp)
	return resp, nil
}
