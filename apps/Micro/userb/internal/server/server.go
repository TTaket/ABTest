package server

import (
	conf "ABTest/apps/Micro/userb/internal/config"
	"ABTest/apps/Micro/userb/internal/service"
	logger "ABTest/pkgs/logger"
	pb "ABTest/pkgs/proto/pb_userb"
	"io"

	xetcd "ABTest/pkgs/xetcd"
	xgrpc "ABTest/pkgs/xgrpc"

	"google.golang.org/grpc"
)

var (
	port = conf.GetConfig().Grpc.Port
)

// server 负责完成逻辑和grpc的融合
type UserbServer interface {
	Run()
}

// NewXxxServer
func NewUserbServer(config *conf.Config, s service.UserbService) UserbServer {
	return &userbServer{cfg: config, userbservice: s}
}

type userbServer struct {
	pb.UnimplementedUserbServiceServer
	cfg          *conf.Config
	userbservice service.UserbService
}

// Run starts the gRPC server
func (s *userbServer) Run() {

	var (
		srv     *grpc.Server
		err     error
		closeio io.Closer
	)

	//配置grpc服务
	srv, closeio, err = xgrpc.NewServer(s.cfg.Grpc)
	if err != nil {
		logger.Error("Failed to create grpc server: %v", err)
	}
	defer func() {
		if closeio != nil {
			closeio.Close()
		}
	}()

	//注册服务
	pb.RegisterUserbServiceServer(srv, s)

	//注册etcd
	if s.cfg.Etcd != nil {
		err := xetcd.RegisterEtcd(s.cfg.Etcd.Schema, s.cfg.Etcd.Endpoints, "0.0.0.0", port, s.cfg.Name, 10)
		if err != nil {
			logger.Error("Failed to register etcd: %v", err)
		}
	}

	//启动grpc服务
	xgrpc.RunServer(srv, *xgrpc.NewGrpcServerConfigs(s.cfg.Grpc))
}
