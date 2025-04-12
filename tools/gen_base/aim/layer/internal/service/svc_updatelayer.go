package service

import (
	"context"

	conf "ABTest/apps/Micro/layer/internal/config"
	pb "ABTest/pkgs/proto/pb_layer"
)

// implement the interface
func (s *layerService) UpdateLayer(ctx context.Context, req *pb.UpdateLayerRequest) (resp *pb.UpdateLayerResponse, err error) {
	conf.Log.Infof("UpdateLayer service begin: %v", req)

	// 这里实现逻辑

	
	conf.Log.Infof("UpdateLayer service end: %v", resp)
	return resp, nil
}
