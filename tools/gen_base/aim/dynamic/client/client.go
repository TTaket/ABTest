package client

import (
	"context"

	conf "ABTest/apps/Micro/dynamic/internal/config"
	log "ABTest/pkgs/logger"
	pb "ABTest/pkgs/proto/pb_dynamic"
	"ABTest/pkgs/xgrpc"

	"github.com/opentracing/opentracing-go"
	"google.golang.org/grpc"
)

type DynamicClient interface {
	CreateLayerUserBinding(ctx context.Context, in *pb.CreateLayerUserBindingRequest, opts ...grpc.CallOption) (*pb.CreateLayerUserBindingResponse, error)
	GetLayerUserBinding(ctx context.Context, in *pb.GetLayerUserBindingRequest, opts ...grpc.CallOption) (*pb.GetLayerUserBindingResponse, error)
	DeleteLayerUserBinding(ctx context.Context, in *pb.DeleteLayerUserBindingRequest, opts ...grpc.CallOption) (*pb.DeleteLayerUserBindingResponse, error)
	CreateExperimentUserBinding(ctx context.Context, in *pb.CreateExperimentUserBindingRequest, opts ...grpc.CallOption) (*pb.CreateExperimentUserBindingResponse, error)
	GetExperimentUserBinding(ctx context.Context, in *pb.GetExperimentUserBindingRequest, opts ...grpc.CallOption) (*pb.GetExperimentUserBindingResponse, error)
	DeleteExperimentUserBinding(ctx context.Context, in *pb.DeleteExperimentUserBindingRequest, opts ...grpc.CallOption) (*pb.DeleteExperimentUserBindingResponse, error)
	CreateExperimentLayerBinding(ctx context.Context, in *pb.CreateExperimentLayerBindingRequest, opts ...grpc.CallOption) (*pb.CreateExperimentLayerBindingResponse, error)
	GetExperimentLayerBinding(ctx context.Context, in *pb.GetExperimentLayerBindingRequest, opts ...grpc.CallOption) (*pb.GetExperimentLayerBindingResponse, error)
	GetLayerExperiments(ctx context.Context, in *pb.GetLayerExperimentsRequest, opts ...grpc.CallOption) (*pb.GetLayerExperimentsResponse, error)
	DeleteExperimentLayerBinding(ctx context.Context, in *pb.DeleteExperimentLayerBindingRequest, opts ...grpc.CallOption) (*pb.DeleteExperimentLayerBindingResponse, error)
	
	Logger() *log.MyLogger
}

func NewDynamicClient() DynamicClient {
	//log
	logger := log.NewMyLogger(conf.GetConfig().Log, conf.GetConfig().Name+"_client")
	//xgrpc
	conn, tracer, err := getClientConn()
	if err != nil || conn == nil {
		logger.Errorf("did not connect: %v", err)
		return nil
	}
	return &dynamicClient{Conn: conn, Tracer: tracer, MyLogger: logger}
}

type dynamicClient struct {
	Conn     *grpc.ClientConn
	Tracer   opentracing.Tracer
	MyLogger *log.MyLogger
}

func (c *dynamicClient) Logger() *log.MyLogger {
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
