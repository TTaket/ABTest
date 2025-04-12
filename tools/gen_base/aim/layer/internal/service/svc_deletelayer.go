package service

import (
	"context"

	conf "ABTest/apps/Micro/layer/internal/config"
	pb "ABTest/pkgs/proto/pb_layer"
)

// implement the interface
func (s *layerService) DeleteLayer(ctx context.Context, req *pb.DeleteLayerRequest) (resp *pb.DeleteLayerResponse, err error) {
	conf.Log.Infof("DeleteLayer service begin: %v", req)

	// 这里实现逻辑

	
	conf.Log.Infof("DeleteLayer service end: %v", resp)
	return resp, nil
}
