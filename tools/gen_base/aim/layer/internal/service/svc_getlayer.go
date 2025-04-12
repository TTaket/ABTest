package service

import (
	"context"

	conf "ABTest/apps/Micro/layer/internal/config"
	pb "ABTest/pkgs/proto/pb_layer"
)

// implement the interface
func (s *layerService) GetLayer(ctx context.Context, req *pb.GetLayerRequest) (resp *pb.GetLayerResponse, err error) {
	conf.Log.Infof("GetLayer service begin: %v", req)

	// 这里实现逻辑

	
	conf.Log.Infof("GetLayer service end: %v", resp)
	return resp, nil
}
