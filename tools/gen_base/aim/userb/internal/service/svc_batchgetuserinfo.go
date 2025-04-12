package service

import (
	"context"

	conf "ABTest/apps/Micro/userb/internal/config"
	pb "ABTest/pkgs/proto/pb_userb"
)

// implement the interface
func (s *userbService) BatchGetUserInfo(ctx context.Context, req *pb.BatchGetUserInfoRequest) (resp *pb.BatchGetUserInfoResponse, err error) {
	conf.Log.Infof("BatchGetUserInfo service begin: %v", req)

	// 这里实现逻辑

	
	conf.Log.Infof("BatchGetUserInfo service end: %v", resp)
	return resp, nil
}
