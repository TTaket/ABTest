package service

import (
	"context"

	conf "ABTest/apps/Micro/dynamic/internal/config"
	pb "ABTest/pkgs/proto/pb_dynamic"
)

// implement the interface
func (s *dynamicService) DeleteLayerUserBinding(ctx context.Context, req *pb.DeleteLayerUserBindingRequest) (resp *pb.DeleteLayerUserBindingResponse, err error) {
	conf.Log.Infof("DeleteLayerUserBinding service begin: %v", req)

	// 这里实现逻辑

	
	conf.Log.Infof("DeleteLayerUserBinding service end: %v", resp)
	return resp, nil
}
