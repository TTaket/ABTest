package service

import (
	"context"

	conf "ABTest/apps/Micro/dynamic/internal/config"
	pb "ABTest/pkgs/proto/pb_dynamic"
)

// implement the interface
func (s *dynamicService) CreateExperimentUserBinding(ctx context.Context, req *pb.CreateExperimentUserBindingRequest) (resp *pb.CreateExperimentUserBindingResponse, err error) {
	conf.Log.Infof("CreateExperimentUserBinding service begin: %v", req)

	// 这里实现逻辑

	
	conf.Log.Infof("CreateExperimentUserBinding service end: %v", resp)
	return resp, nil
}
