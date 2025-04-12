package service

import (
	"context"

	conf "ABTest/apps/Micro/layer/internal/config"
	pb "ABTest/pkgs/proto/pb_layer"
)

// implement the interface
func (s *layerService) CreateLayer(ctx context.Context, req *pb.CreateLayerRequest) (resp *pb.CreateLayerResponse, err error) {
	conf.Log.Infof("CreateLayer service begin: %v", req)

	// 这里实现逻辑

	
	conf.Log.Infof("CreateLayer service end: %v", resp)
	return resp, nil
}
