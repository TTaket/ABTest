package service

import (
	"context"

	conf "ABTest/apps/Micro/userb/internal/config"
	pb "ABTest/pkgs/proto/pb_userb"
)

// implement the interface
func (s *userbService) AddUser(ctx context.Context, req *pb.AddUserRequest) (resp *pb.AddUserResponse, err error) {
	conf.Log.Infof("AddUser service begin: %v", req)

	// 这里实现逻辑

	
	conf.Log.Infof("AddUser service end: %v", resp)
	return resp, nil
}
