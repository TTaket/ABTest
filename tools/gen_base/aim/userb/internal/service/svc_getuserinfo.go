package service

import (
	"context"

	conf "ABTest/apps/Micro/userb/internal/config"
	pb "ABTest/pkgs/proto/pb_userb"
)

// implement the interface
func (s *userbService) GetUserInfo(ctx context.Context, req *pb.GetUserInfoRequest) (resp *pb.GetUserInfoResponse, err error) {
	conf.Log.Infof("GetUserInfo service begin: %v", req)

	// 这里实现逻辑

	
	conf.Log.Infof("GetUserInfo service end: %v", resp)
	return resp, nil
}
