package client

import (
	"context"

	conf "ABTest/apps/Micro/%1/internal/config"
	log "ABTest/pkgs/logger"
	pb "ABTest/pkgs/proto/pb_%1"
	"ABTest/pkgs/xgrpc"

	"github.com/opentracing/opentracing-go"
	"google.golang.org/grpc"
)

type %2Client interface {
	%00001
	Logger() *log.MyLogger
}

func New%2Client() %2Client {
	//log
	logger := log.NewMyLogger(conf.GetConfig().Log, conf.GetConfig().Name+"_client")
	//xgrpc
	conn, tracer, err := getClientConn()
	if err != nil || conn == nil {
		logger.Errorf("did not connect: %v", err)
		return nil
	}
	return &%3Client{Conn: conn, Tracer: tracer, MyLogger: logger}
}

type %3Client struct {
	Conn     *grpc.ClientConn
	Tracer   opentracing.Tracer
	MyLogger *log.MyLogger
}

func (c *%3Client) Logger() *log.MyLogger {
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
