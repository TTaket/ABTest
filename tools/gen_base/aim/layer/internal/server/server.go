package server

import (
	conf "ABTest/apps/Micro/layer/internal/config"
	"ABTest/apps/Micro/layer/internal/service"
	pb "ABTest/pkgs/proto/pb_layer"
	"io"

	xetcd "ABTest/pkgs/xetcd"
	xgrpc "ABTest/pkgs/xgrpc"

	"google.golang.org/grpc"
)

var (
	port = conf.GetConfig().Grpc.Port
)

// server 负责完成逻辑和grpc的融合
type LayerServer interface {
	Run()
}

// NewLayerServer returns a new LayerServer
func NewLayerServer(config *conf.Config, s service.LayerService) LayerServer {
	return &layerServer{cfg: config, layerService: s}
}

type layerServer struct {
	pb.UnimplementedLayerServiceServer
	cfg               *conf.Config
	layerService service.LayerService
}

// Run starts the gRPC server
func (s *layerServer) Run() {

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
	pb.RegisterLayerServiceServer(srv, s)

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
