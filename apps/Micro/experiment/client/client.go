package client

import (
	"context"

	conf "ABTest/apps/Micro/experiment/internal/config"
	log "ABTest/pkgs/logger"
	pb "ABTest/pkgs/proto/pb_experiment"
	"ABTest/pkgs/xgrpc"

	"github.com/opentracing/opentracing-go"
	"google.golang.org/grpc"
)

// TODO 同一个客户端 多次调用GetExperiment和CreateExperiment 会多次的调用connect
type ExperimentClient interface {
	GetExperiment(ctx context.Context, in *pb.GetExperimentRequest, opts ...grpc.CallOption) (*pb.GetExperimentResponse, error)
	CreateExperiment(ctx context.Context, in *pb.CreateExperimentRequest, opts ...grpc.CallOption) (*pb.CreateExperimentResponse, error)
}

func NewExperimentClient() ExperimentClient {
	conn, tracer, err := getClientConn()
	if err != nil || conn == nil {
		log.Errorf("did not connect: %v", err)
		return nil
	}
	return &experimentClient{Conn: conn, Tracer: tracer}
}

type experimentClient struct {
	Conn   *grpc.ClientConn
	Tracer opentracing.Tracer
}

func getClientConn() (*grpc.ClientConn, opentracing.Tracer, error) {
	// 填装Client配置
	configs, tracer, _ := xgrpc.NewGrpcClientConfigs(conf.GetConfig().Name, conf.GetConfig().Grpc, conf.GetConfig().Etcd, conf.GetConfig().Grpc.Jaeger)

	// 获取连接并返回
	conn, err := xgrpc.GetClientConn(configs)
	return conn, tracer, err
}
