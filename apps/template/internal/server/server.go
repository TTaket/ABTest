package server

import (
	conf "ABTest/apps/Xxx/internal/config"
	"ABTest/apps/Xxx/internal/service"
	logger "ABTest/pkgs/logger"
	pb "ABTest/pkgs/proto/pb_Xxx"
	"io"

	xetcd "ABTest/pkgs/xetcd"
	xgrpc "ABTest/pkgs/xgrpc"

	"google.golang.org/grpc"
)

var (
	port = conf.GetConfig().Grpc.Port
)

// server 负责完成逻辑和grpc的融合
type XxxServer interface {
	Run()
}

// NewXxxServer
func NewXxxServer(config *conf.Config, s service.ExperimentService) XxxServer {
	return &xxxServer{cfg: config, xxxService: s}
}

type xxxServer struct {
	pb.UnimplementedExperimentServiceServer
	cfg *conf.Config
	// TODO: 请根据实际情况修改xxxService的类型
	xxxService service.xxxService
}

// Run starts the gRPC server
func (s *xxxServer) Run() {

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
	pb.RegisterXxxServiceServer(srv, s)

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
