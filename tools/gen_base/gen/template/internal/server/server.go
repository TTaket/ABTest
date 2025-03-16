package server

import (
	conf "ABTest/apps/Micro/%1/internal/config"
	"ABTest/apps/Micro/%1/internal/service"
	pb "ABTest/pkgs/proto/pb_%1"
	"io"

	xetcd "ABTest/pkgs/xetcd"
	xgrpc "ABTest/pkgs/xgrpc"

	"google.golang.org/grpc"
)

var (
	port = conf.GetConfig().Grpc.Port
)

// server 负责完成逻辑和grpc的融合
type %2Server interface {
	Run()
}

// New%2Server returns a new %2Server
func New%2Server(config *conf.Config, s service.%2Service) %2Server {
	return &%3Server{cfg: config, %3Service: s}
}

type %3Server struct {
	pb.Unimplemented%4Server
	cfg               *conf.Config
	%3Service service.%2Service
}

// Run starts the gRPC server
func (s *%3Server) Run() {

	var (
		srv     *grpc.Server
		err     error
		closeio io.Closer
	)

	//配置grpc服务
	srv, closeio, err = xgrpc.NewServer(s.cfg.Grpc)
	if err != nil {
		conf.Log.Errorf("Failed to create grpc server: %v", err)
	}
	defer func() {
		if closeio != nil {
			closeio.Close()
		}
	}()
	conf.Log.Infof("grpc Create success")

	//注册服务
	pb.Register%4Server(srv, s)

	//注册etcd
	if s.cfg.Etcd != nil {
		err := xetcd.RegisterEtcd(s.cfg.Etcd.Schema, s.cfg.Etcd.Endpoints, "0.0.0.0", port, s.cfg.Name, 10)
		if err != nil {
			conf.Log.Errorf("Failed to register etcd: %v", err)
		}
		conf.Log.Infof("etcd Register success")
	}

	//启动grpc服务
	xgrpc.RunServer(srv, *xgrpc.NewGrpcServerConfigs(s.cfg.Grpc))
}
