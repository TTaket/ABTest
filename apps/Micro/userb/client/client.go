package client

import (
	log "ABTest/pkgs/logger"
	"ABTest/pkgs/xgrpc"
	"context"

	pb "ABTest/pkgs/proto/pb_userb"

	"github.com/opentracing/opentracing-go"
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
	conn, tracer, err := getClientConn()
	if err != nil || conn == nil {
		log.Errorf("did not connect: %v", err)
		return nil
	}
	return &userbClient{conn: conn, tracer: tracer}
}

type userbClient struct {
	conn   *grpc.ClientConn
	tracer opentracing.Tracer
}

// 内部使用
func getClientConn() (*grpc.ClientConn, opentracing.Tracer, error) {
	// 填装Client配置
	configs, tracer, _ := xgrpc.NewGrpcClientConfigs(conf.GetConfig().Name, conf.GetConfig().Grpc, conf.GetConfig().Etcd, conf.GetConfig().Grpc.Jaeger)

	// 获取连接并返回
	conn, err := xgrpc.GetClientConn(configs)
	return conn, tracer, err
}
