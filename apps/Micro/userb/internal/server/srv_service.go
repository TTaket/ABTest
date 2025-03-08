package server

import (
	pb "ABTest/pkgs/proto/pb_userb"
	"context"
)

func (s *userbServer) AddUser(ctx context.Context, in *pb.AddUserRequest) (*pb.AddUserResponse, error) {
	return s.userbservice.AddUser(ctx, in)
}

func (s *userbServer) BatchAddUser(ctx context.Context, in *pb.BatchAddUserRequest) (*pb.BatchAddUserResponse, error) {
	return s.userbservice.BatchAddUser(ctx, in)
}

func (s *userbServer) DeleteUser(ctx context.Context, in *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error) {
	return s.userbservice.DeleteUser(ctx, in)
}

func (s *userbServer) BatchDeleteUser(ctx context.Context, in *pb.BatchDeleteUserRequest) (*pb.BatchDeleteUserResponse, error) {
	return s.userbservice.BatchDeleteUser(ctx, in)
}

func (s *userbServer) UpdateUser(ctx context.Context, in *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error) {
	return s.userbservice.UpdateUser(ctx, in)
}

func (s *userbServer) BatchUpdateUser(ctx context.Context, in *pb.BatchUpdateUserRequest) (*pb.BatchUpdateUserResponse, error) {
	return s.userbservice.BatchUpdateUser(ctx, in)
}

func (s *userbServer) GetUserInfo(ctx context.Context, in *pb.GetUserInfoRequest) (*pb.GetUserInfoResponse, error) {
	return s.userbservice.GetUserInfo(ctx, in)
}

func (s *userbServer) BatchGetUserInfo(ctx context.Context, in *pb.BatchGetUserInfoRequest) (*pb.BatchGetUserInfoResponse, error) {
	return s.userbservice.BatchGetUserInfo(ctx, in)
}
