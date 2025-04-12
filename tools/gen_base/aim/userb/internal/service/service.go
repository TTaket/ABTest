package service

import (
	"context"

	pb "ABTest/pkgs/proto/pb_userb"
		
	"google.golang.org/grpc"
)

// service 负责业务逻辑实现
type UserbService interface {
	AddUser(ctx context.Context, in *pb.AddUserRequest) (*pb.AddUserResponse, error)
	BatchAddUser(ctx context.Context, in *pb.BatchAddUserRequest) (*pb.BatchAddUserResponse, error)
	DeleteUser(ctx context.Context, in *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error)
	BatchDeleteUser(ctx context.Context, in *pb.BatchDeleteUserRequest) (*pb.BatchDeleteUserResponse, error)
	UpdateUser(ctx context.Context, in *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error)
	BatchUpdateUser(ctx context.Context, in *pb.BatchUpdateUserRequest) (*pb.BatchUpdateUserResponse, error)
	GetUserInfo(ctx context.Context, in *pb.GetUserInfoRequest) (*pb.GetUserInfoResponse, error)
	BatchGetUserInfo(ctx context.Context, in *pb.BatchGetUserInfoRequest) (*pb.BatchGetUserInfoResponse, error)
	ScatterTraffic(ctx context.Context, in *pb.ScatterTrafficRequest) (*pb.ScatterTrafficResponse, error)
	
}

type userbService struct {

	// others..
}

func NewUserbService() UserbService {
	return &userbService{}
}
