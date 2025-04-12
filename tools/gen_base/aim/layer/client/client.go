package client

import (
	"context"

	conf "ABTest/apps/Micro/layer/internal/config"
	log "ABTest/pkgs/logger"
	pb "ABTest/pkgs/proto/pb_layer"
	"ABTest/pkgs/xgrpc"

	"github.com/opentracing/opentracing-go"
	"google.golang.org/grpc"
)

type LayerClient interface {
	CreateLayer(ctx context.Context, in *pb.CreateLayerRequest, opts ...grpc.CallOption) (*pb.CreateLayerResponse, error)
	GetLayer(ctx context.Context, in *pb.GetLayerRequest, opts ...grpc.CallOption) (*pb.GetLayerResponse, error)
	UpdateLayer(ctx context.Context, in *pb.UpdateLayerRequest, opts ...grpc.CallOption) (*pb.UpdateLayerResponse, error)
	DeleteLayer(ctx context.Context, in *pb.DeleteLayerRequest, opts ...grpc.CallOption) (*pb.DeleteLayerResponse, error)
	ListLayers(ctx context.Context, in *pb.ListLayersRequest, opts ...grpc.CallOption) (*pb.ListLayersResponse, error)
	
	Logger() *log.MyLogger
}

func NewLayerClient() LayerClient {
	//log
	logger := log.NewMyLogger(conf.GetConfig().Log, conf.GetConfig().Name+"_client")
	//xgrpc
	conn, tracer, err := getClientConn()
	if err != nil || conn == nil {
		logger.Errorf("did not connect: %v", err)
		return nil
	}
	return &layerClient{Conn: conn, Tracer: tracer, MyLogger: logger}
}

type layerClient struct {
	Conn     *grpc.ClientConn
	Tracer   opentracing.Tracer
	MyLogger *log.MyLogger
}

func (c *layerClient) Logger() *log.MyLogger {
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
