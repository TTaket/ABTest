package service

import (
	"context"

	conf "ABTest/apps/Micro/dynamic/internal/config"
	pb "ABTest/pkgs/proto/pb_dynamic"
)

// implement the interface
func (s *dynamicService) DeleteExperimentLayerBinding(ctx context.Context, req *pb.DeleteExperimentLayerBindingRequest) (resp *pb.DeleteExperimentLayerBindingResponse, err error) {
	conf.Log.Infof("DeleteExperimentLayerBinding service begin: %v", req)

	// 这里实现逻辑

	
	conf.Log.Infof("DeleteExperimentLayerBinding service end: %v", resp)
	return resp, nil
}
