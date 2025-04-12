package service

import (
	"context"

	conf "ABTest/apps/Micro/userb/internal/config"
	pb "ABTest/pkgs/proto/pb_userb"
)

// implement the interface
func (s *userbService) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (resp *pb.DeleteUserResponse, err error) {
	conf.Log.Infof("DeleteUser service begin: %v", req)

	// 这里实现逻辑

	conf.Log.Infof("DeleteUser service end: %v", resp)
	return resp, nil
}
