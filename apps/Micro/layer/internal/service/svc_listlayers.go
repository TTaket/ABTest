package service

import (
	"context"

	conf "ABTest/apps/Micro/layer/internal/config"
	pb "ABTest/pkgs/proto/pb_layer"
)

// implement the interface
func (s *layerService) ListLayers(ctx context.Context, req *pb.ListLayersRequest) (resp *pb.ListLayersResponse, err error) {
	conf.Log.Infof("ListLayers service begin: %v", req)

	// 这里实现逻辑

	
	conf.Log.Infof("ListLayers service end: %v", resp)
	return resp, nil
}
