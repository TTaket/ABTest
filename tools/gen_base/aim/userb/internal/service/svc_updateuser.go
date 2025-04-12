package service

import (
	"context"

	conf "ABTest/apps/Micro/userb/internal/config"
	pb "ABTest/pkgs/proto/pb_userb"
)

// implement the interface
func (s *userbService) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (resp *pb.UpdateUserResponse, err error) {
	conf.Log.Infof("UpdateUser service begin: %v", req)

	// 这里实现逻辑

	
	conf.Log.Infof("UpdateUser service end: %v", resp)
	return resp, nil
}
