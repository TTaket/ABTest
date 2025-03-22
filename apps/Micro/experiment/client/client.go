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

type ExperimentClient interface {
	CreateExperiment(ctx context.Context, in *pb.CreateExperimentRequest, opts ...grpc.CallOption) (*pb.CreateExperimentResponse, error)
	GetExperiment(ctx context.Context, in *pb.GetExperimentRequest, opts ...grpc.CallOption) (*pb.GetExperimentResponse, error)
	DeleteExperiment(ctx context.Context, in *pb.DeleteExperimentRequest, opts ...grpc.CallOption) (*pb.DeleteExperimentResponse, error)
	UpdateExperimentBaseInfo(ctx context.Context, in *pb.UpdateExperimentBaseInfoRequest, opts ...grpc.CallOption) (*pb.UpdateExperimentBaseInfoResponse, error)
	AddExperimentGroup(ctx context.Context, in *pb.AddExperimentGroupRequest, opts ...grpc.CallOption) (*pb.AddExperimentGroupResponse, error)
	DeleteExperimentGroup(ctx context.Context, in *pb.DeleteExperimentGroupRequest, opts ...grpc.CallOption) (*pb.DeleteExperimentGroupResponse, error)
	
	Logger() *log.MyLogger
}

func NewExperimentClient() ExperimentClient {
	//log
	logger := log.NewMyLogger(conf.GetConfig().Log, conf.GetConfig().Name+"_client")
	//xgrpc
	conn, tracer, err := getClientConn()
	if err != nil || conn == nil {
		logger.Errorf("did not connect: %v", err)
		return nil
	}
	return &experimentClient{Conn: conn, Tracer: tracer, MyLogger: logger}
}

type experimentClient struct {
	Conn     *grpc.ClientConn
	Tracer   opentracing.Tracer
	MyLogger *log.MyLogger
}

func (c *experimentClient) Logger() *log.MyLogger {
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
