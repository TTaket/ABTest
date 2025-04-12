package service

import (
	"context"

	conf "ABTest/apps/Micro/userb/internal/config"
	pb "ABTest/pkgs/proto/pb_userb"
)

// implement the interface
func (s *userbService) ScatterTraffic(ctx context.Context, req *pb.ScatterTrafficRequest) (resp *pb.ScatterTrafficResponse, err error) {
	conf.Log.Infof("ScatterTraffic service begin: %v", req)

	// 这里实现逻辑

	
	conf.Log.Infof("ScatterTraffic service end: %v", resp)
	return resp, nil
}
