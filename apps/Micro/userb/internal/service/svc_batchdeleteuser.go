package service

import (
	"context"

	conf "ABTest/apps/Micro/userb/internal/config"
	pb "ABTest/pkgs/proto/pb_userb"
)

// implement the interface
func (s *userbService) BatchDeleteUser(ctx context.Context, req *pb.BatchDeleteUserRequest) (resp *pb.BatchDeleteUserResponse, err error) {
	conf.Log.Infof("BatchDeleteUser service begin: %v", req)

	// 这里实现逻辑

	conf.Log.Infof("BatchDeleteUser service end: %v", resp)
	return resp, nil
}
