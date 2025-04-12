package service

import (
	"context"

	conf "ABTest/apps/Micro/userb/internal/config"
	pb "ABTest/pkgs/proto/pb_userb"
)

// implement the interface
func (s *userbService) BatchUpdateUser(ctx context.Context, req *pb.BatchUpdateUserRequest) (resp *pb.BatchUpdateUserResponse, err error) {
	conf.Log.Infof("BatchUpdateUser service begin: %v", req)

	// 这里实现逻辑

	
	conf.Log.Infof("BatchUpdateUser service end: %v", resp)
	return resp, nil
}
