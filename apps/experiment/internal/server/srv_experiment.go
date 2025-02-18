package server

import (
	conf "ABTest/apps/experiment/internal/config"
	"ABTest/apps/experiment/internal/service"
	logger "ABTest/pkgs/logger"
	pb "ABTest/pkgs/proto/pb_experiment"
	"io"

	xgrpc "ABTest/pkgs/xgrpc"

	"google.golang.org/grpc"
)

var (
	port = conf.GetConfig().Grpc.Port
)

// server 负责完成逻辑和grpc的融合
type ExperimentServer interface {
	Run()
}

// NewExperimentServer returns a new ExperimentServer
func NewExperimentServer(config *conf.Config, s service.ExperimentService) ExperimentServer {
	return &experimentServer{cfg: config, experimentService: s}
}

type experimentServer struct {
	pb.UnimplementedExperimentServiceServer
	cfg               *conf.Config
	experimentService service.ExperimentService
}

// Run starts the gRPC server
func (s *experimentServer) Run() {

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
	pb.RegisterExperimentServiceServer(srv, s)

	//启动grpc服务
	xgrpc.RunServer(srv, *xgrpc.NewGrpcServerConfigs(s.cfg.Grpc))
}
