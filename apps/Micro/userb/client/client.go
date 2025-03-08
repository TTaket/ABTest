package client

import (
	"ABTest/pkgs/xgrpc"
	"context"

	pb "ABTest/pkgs/proto/pb_userb"

	"google.golang.org/grpc"

	conf "ABTest/apps/Micro/userb/internal/config"
)

type UserbClient interface {
	AddUser(ctx context.Context, in *pb.AddUserRequest) (*pb.AddUserResponse, error)
	BatchAddUser(ctx context.Context, in *pb.BatchAddUserRequest) (*pb.BatchAddUserResponse, error)
	DeleteUser(ctx context.Context, in *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error)
	BatchDeleteUser(ctx context.Context, in *pb.BatchDeleteUserRequest) (*pb.BatchDeleteUserResponse, error)
	UpdateUser(ctx context.Context, in *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error)
	BatchUpdateUser(ctx context.Context, in *pb.BatchUpdateUserRequest) (*pb.BatchUpdateUserResponse, error)
	GetUserInfo(ctx context.Context, in *pb.GetUserInfoRequest) (*pb.GetUserInfoResponse, error)
	BatchGetUserInfo(ctx context.Context, in *pb.BatchGetUserInfoRequest) (*pb.BatchGetUserInfoResponse, error)
}

func NewUserbClient() UserbClient {
	return &userbClient{}
}

type userbClient struct {
}

// 内部使用
func getClientConn() (*grpc.ClientConn, error) {
	// 填装Client配置
	configs, iocloser := xgrpc.NewGrpcClientConfigs(conf.GetConfig().Name, conf.GetConfig().Grpc, conf.GetConfig().Etcd, conf.GetConfig().Grpc.Jaeger)
	defer iocloser.Close()

	// 获取连接并返回
	return xgrpc.GetClientConn(configs)
}
