package service

import (
	"context"

	conf "ABTest/apps/Micro/dynamic/internal/config"
	pb "ABTest/pkgs/proto/pb_dynamic"
)

// implement the interface
func (s *dynamicService) GetExperimentUserBinding(ctx context.Context, req *pb.GetExperimentUserBindingRequest) (resp *pb.GetExperimentUserBindingResponse, err error) {
	conf.Log.Infof("GetExperimentUserBinding service begin: %v", req)

	// 这里实现逻辑

	
	conf.Log.Infof("GetExperimentUserBinding service end: %v", resp)
	return resp, nil
}
