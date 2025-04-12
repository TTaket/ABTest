package service

import (
	"context"

	conf "ABTest/apps/Micro/dynamic/internal/config"
	pb "ABTest/pkgs/proto/pb_dynamic"
)

// implement the interface
func (s *dynamicService) DeleteExperimentUserBinding(ctx context.Context, req *pb.DeleteExperimentUserBindingRequest) (resp *pb.DeleteExperimentUserBindingResponse, err error) {
	conf.Log.Infof("DeleteExperimentUserBinding service begin: %v", req)

	// 这里实现逻辑

	
	conf.Log.Infof("DeleteExperimentUserBinding service end: %v", resp)
	return resp, nil
}
