package server

import (
	pb "ABTest/pkgs/proto/pb_userb"
	"context"
)

// service 负责业务逻辑实现

func (s *userbServer) AddUser(ctx context.Context, req *pb.AddUserRequest) (resp *pb.AddUserResponse, err error) {
	return s.userbService.AddUser(ctx, req)
}

	
func (s *userbServer) BatchAddUser(ctx context.Context, req *pb.BatchAddUserRequest) (resp *pb.BatchAddUserResponse, err error) {
	return s.userbService.BatchAddUser(ctx, req)
}

	
func (s *userbServer) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (resp *pb.DeleteUserResponse, err error) {
	return s.userbService.DeleteUser(ctx, req)
}

	
func (s *userbServer) BatchDeleteUser(ctx context.Context, req *pb.BatchDeleteUserRequest) (resp *pb.BatchDeleteUserResponse, err error) {
	return s.userbService.BatchDeleteUser(ctx, req)
}

	
func (s *userbServer) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (resp *pb.UpdateUserResponse, err error) {
	return s.userbService.UpdateUser(ctx, req)
}

	
func (s *userbServer) BatchUpdateUser(ctx context.Context, req *pb.BatchUpdateUserRequest) (resp *pb.BatchUpdateUserResponse, err error) {
	return s.userbService.BatchUpdateUser(ctx, req)
}

	
func (s *userbServer) GetUserInfo(ctx context.Context, req *pb.GetUserInfoRequest) (resp *pb.GetUserInfoResponse, err error) {
	return s.userbService.GetUserInfo(ctx, req)
}

	
func (s *userbServer) BatchGetUserInfo(ctx context.Context, req *pb.BatchGetUserInfoRequest) (resp *pb.BatchGetUserInfoResponse, err error) {
	return s.userbService.BatchGetUserInfo(ctx, req)
}

	
func (s *userbServer) ScatterTraffic(ctx context.Context, req *pb.ScatterTrafficRequest) (resp *pb.ScatterTrafficResponse, err error) {
	return s.userbService.ScatterTraffic(ctx, req)
}

	
