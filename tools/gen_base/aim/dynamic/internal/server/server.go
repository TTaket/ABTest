package server

import (
	conf "ABTest/apps/Micro/dynamic/internal/config"
	"ABTest/apps/Micro/dynamic/internal/service"
	pb "ABTest/pkgs/proto/pb_dynamic"
	"io"

	xetcd "ABTest/pkgs/xetcd"
	xgrpc "ABTest/pkgs/xgrpc"

	"google.golang.org/grpc"
)

var (
	port = conf.GetConfig().Grpc.Port
)

// server 负责完成逻辑和grpc的融合
type DynamicServer interface {
	Run()
}

// NewDynamicServer returns a new DynamicServer
func NewDynamicServer(config *conf.Config, s service.DynamicService) DynamicServer {
	return &dynamicServer{cfg: config, dynamicService: s}
}

type dynamicServer struct {
	pb.UnimplementedDynamicServiceServer
	cfg               *conf.Config
	dynamicService service.DynamicService
}

// Run starts the gRPC server
func (s *dynamicServer) Run() {

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
	pb.RegisterDynamicServiceServer(srv, s)

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
