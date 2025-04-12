package service

import (
	"context"

	conf "ABTest/apps/Micro/userb/internal/config"
	pb "ABTest/pkgs/proto/pb_userb"
)

// implement the interface
func (s *userbService) BatchAddUser(ctx context.Context, req *pb.BatchAddUserRequest) (resp *pb.BatchAddUserResponse, err error) {
	conf.Log.Infof("BatchAddUser service begin: %v", req)

	// 这里实现逻辑

	
	conf.Log.Infof("BatchAddUser service end: %v", resp)
	return resp, nil
}
