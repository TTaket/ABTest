package client

import (
	"context"

	conf "ABTest/apps/Micro/experiment/internal/config"
	pb "ABTest/pkgs/proto/pb_experiment"
	"ABTest/pkgs/xgrpc"

	"google.golang.org/grpc"
)

// TODO 同一个客户端 多次调用GetExperiment和CreateExperiment 会多次的调用connect
type ExperimentServiceClient interface {
	GetExperiment(ctx context.Context, in *pb.GetExperimentRequest, opts ...grpc.CallOption) (*pb.GetExperimentResponse, error)
	CreateExperiment(ctx context.Context, in *pb.CreateExperimentRequest, opts ...grpc.CallOption) (*pb.CreateExperimentResponse, error)
}

func NewExperimentServiceClient() ExperimentServiceClient {
	return &experimentServiceClient{}
}

type experimentServiceClient struct {
}

func getClientConn() (*grpc.ClientConn, error) {
	configs, iocloser := xgrpc.NewGrpcClientConfigs(conf.GetConfig().Name, conf.GetConfig().Grpc, conf.GetConfig().Etcd, conf.GetConfig().Grpc.Jaeger)
	defer iocloser.Close()
	return xgrpc.GetClientConn(configs)
}
