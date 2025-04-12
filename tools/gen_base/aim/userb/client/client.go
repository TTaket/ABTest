package client

import (
	"context"

	conf "ABTest/apps/Micro/userb/internal/config"
	log "ABTest/pkgs/logger"
	pb "ABTest/pkgs/proto/pb_userb"
	"ABTest/pkgs/xgrpc"

	"github.com/opentracing/opentracing-go"
	"google.golang.org/grpc"
)

type UserbClient interface {
	AddUser(ctx context.Context, in *pb.AddUserRequest, opts ...grpc.CallOption) (*pb.AddUserResponse, error)
	BatchAddUser(ctx context.Context, in *pb.BatchAddUserRequest, opts ...grpc.CallOption) (*pb.BatchAddUserResponse, error)
	DeleteUser(ctx context.Context, in *pb.DeleteUserRequest, opts ...grpc.CallOption) (*pb.DeleteUserResponse, error)
	BatchDeleteUser(ctx context.Context, in *pb.BatchDeleteUserRequest, opts ...grpc.CallOption) (*pb.BatchDeleteUserResponse, error)
	UpdateUser(ctx context.Context, in *pb.UpdateUserRequest, opts ...grpc.CallOption) (*pb.UpdateUserResponse, error)
	BatchUpdateUser(ctx context.Context, in *pb.BatchUpdateUserRequest, opts ...grpc.CallOption) (*pb.BatchUpdateUserResponse, error)
	GetUserInfo(ctx context.Context, in *pb.GetUserInfoRequest, opts ...grpc.CallOption) (*pb.GetUserInfoResponse, error)
	BatchGetUserInfo(ctx context.Context, in *pb.BatchGetUserInfoRequest, opts ...grpc.CallOption) (*pb.BatchGetUserInfoResponse, error)
	ScatterTraffic(ctx context.Context, in *pb.ScatterTrafficRequest, opts ...grpc.CallOption) (*pb.ScatterTrafficResponse, error)
	
	Logger() *log.MyLogger
}

func NewUserbClient() UserbClient {
	//log
	logger := log.NewMyLogger(conf.GetConfig().Log, conf.GetConfig().Name+"_client")
	//xgrpc
	conn, tracer, err := getClientConn()
	if err != nil || conn == nil {
		logger.Errorf("did not connect: %v", err)
		return nil
	}
	return &userbClient{Conn: conn, Tracer: tracer, MyLogger: logger}
}

type userbClient struct {
	Conn     *grpc.ClientConn
	Tracer   opentracing.Tracer
	MyLogger *log.MyLogger
}

func (c *userbClient) Logger() *log.MyLogger {
	return c.MyLogger
}

func getClientConn() (*grpc.ClientConn, opentracing.Tracer, error) {
	// 填装Client配置
	configs, tracer, _ := xgrpc.NewGrpcClientConfigs(conf.GetConfig().Name, conf.GetConfig().Grpc, conf.GetConfig().Etcd, conf.GetConfig().Grpc.Jaeger)

	// 获取连接并返回
	conn, err := xgrpc.GetClientConn(configs)
	if err != nil {
		panic(err)
	}
	return conn, tracer, err
}
