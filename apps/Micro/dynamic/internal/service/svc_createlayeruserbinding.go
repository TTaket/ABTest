package service

import (
	"context"

	conf "ABTest/apps/Micro/dynamic/internal/config"
	pb "ABTest/pkgs/proto/pb_dynamic"
)

// implement the interface
func (s *dynamicService) CreateLayerUserBinding(ctx context.Context, req *pb.CreateLayerUserBindingRequest) (resp *pb.CreateLayerUserBindingResponse, err error) {
	conf.Log.Infof("CreateLayerUserBinding service begin: %v", req)

	// 这里实现逻辑

	
	conf.Log.Infof("CreateLayerUserBinding service end: %v", resp)
	return resp, nil
}
