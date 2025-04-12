package service

import (
	"context"

	conf "ABTest/apps/Micro/dynamic/internal/config"
	pb "ABTest/pkgs/proto/pb_dynamic"
)

// implement the interface
func (s *dynamicService) GetLayerUserBinding(ctx context.Context, req *pb.GetLayerUserBindingRequest) (resp *pb.GetLayerUserBindingResponse, err error) {
	conf.Log.Infof("GetLayerUserBinding service begin: %v", req)

	// 这里实现逻辑

	
	conf.Log.Infof("GetLayerUserBinding service end: %v", resp)
	return resp, nil
}
