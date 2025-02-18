package xgrpc

import (
	"io"
	"net"
	"strconv"
	"time"

	conf "ABTest/pkgs/config"
	logger "ABTest/pkgs/logger"
	Jaeger "ABTest/pkgs/tracer"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	otgrpc "github.com/opentracing-contrib/go-grpc"
	"github.com/opentracing/opentracing-go"
	"golang.org/x/net/netutil"
	grpc "google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
)

type GrpcServerConfigs struct {
	grpc *conf.Grpc
}

func NewGrpcServerConfigs(grpc *conf.Grpc) *GrpcServerConfigs {
	return &GrpcServerConfigs{grpc}
}

func NewServer(c *conf.Grpc) (grpcServer *grpc.Server, closer io.Closer, err error) {
	var opts []grpc.ServerOption
	var unaryInterceptors []grpc.UnaryServerInterceptor

	// Add recovery interceptor
	unaryInterceptors = append(unaryInterceptors, grpc_recovery.UnaryServerInterceptor())

	// Add KeepParams
	keepParams := grpc.KeepaliveParams(keepalive.ServerParameters{
		MaxConnectionIdle:     time.Duration(c.MaxConnectionIdle) * time.Millisecond,
		MaxConnectionAge:      time.Duration(c.MaxConnectionAge) * time.Millisecond,
		MaxConnectionAgeGrace: time.Duration(c.MaxConnectionAgeGrace) * time.Millisecond,
		Time:                  time.Duration(c.Time) * time.Millisecond,
		Timeout:               time.Duration(c.Timeout) * time.Millisecond,
	})
	opts = append(opts, keepParams)
	// Add Jaeger tracer
	if c.Jaeger.Enabled {
		var tracer opentracing.Tracer
		tracer, closer, err = Jaeger.NewTracer(c.Name, c.Jaeger)
		if err == nil {
			unaryInterceptors = append(unaryInterceptors, otgrpc.OpenTracingServerInterceptor(tracer))
		} else {
			logger.Errorf("Failed to create Jaeger tracer: %v", err)
			return grpcServer, closer, err
		}
	}

	if c.StreamsLimit > 0 {
		// 一个连接中最大并发Stream数
		opts = append(opts, grpc.MaxConcurrentStreams(c.StreamsLimit))
	}
	if c.MaxRecvMsgSize > 0 {
		// 允许接收的最大消息长度
		opts = append(opts, grpc.MaxRecvMsgSize(c.MaxRecvMsgSize))
	}
	if len(unaryInterceptors) > 0 {
		opts = append(opts, grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(unaryInterceptors...)))
	}

	grpcServer = grpc.NewServer(opts...)
	return
}

func RunServer(s *grpc.Server, c GrpcServerConfigs) error {
	var (
		address  string
		listener net.Listener
		err      error
	)
	defer func() {
		s.GracefulStop()
	}()

	address = "0.0.0.0:" + strconv.Itoa(c.grpc.Port)
	listener, err = net.Listen("tcp", address)
	if err != nil {
		logger.Error(err.Error())
		return err
	}
	if c.grpc.ConnectionLimit > 0 {
		// 最大并发连接数
		listener = netutil.LimitListener(listener, c.grpc.ConnectionLimit)
	}

	err = s.Serve(listener)
	if err != nil {
		logger.Error(err.Error())
		return err
	}
	return nil
}
