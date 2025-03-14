package server2

import (
	"context"
	"fmt"
	tracer "grpcdemo/pkg/tracer"
	pb_server2 "grpcdemo/proto/server2/pb_server2/pb_server2"
	"io"
	"log"
	"net"

	grpc_opentracing "github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"
	"github.com/opentracing/opentracing-go"
	"google.golang.org/grpc"
)

// 直接实现 pb_server2.Server1ServiceServer 接口
type server2Service struct {
	pb_server2.UnimplementedServer2ServiceServer // 嵌入 UnimplementedServer2ServiceServer
}

func setTracer_server2() io.Closer {
	trace, iocloser, err := tracer.InitJaeger("server2")
	if err != nil {
		log.Fatalf("Could not initialize jaeger tracer: %s", err.Error())
	}
	opentracing.SetGlobalTracer(trace)
	return iocloser
}

func (s *server2Service) Get(ctx context.Context, req *pb_server2.GetRequest) (*pb_server2.GetResponse, error) {
	resp := &pb_server2.GetResponse{
		ExperimentId: req.GetExperimentId(),
	}
	return resp, nil
}

func StartServer() {
	iocloser := setTracer_server2()
	defer iocloser.Close()
	// Start the server
	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer(
		grpc.UnaryInterceptor(
			grpc_opentracing.UnaryServerInterceptor(),
		),
	)
	ss := &server2Service{}                        // 直接创建 server1Service 实例
	pb_server2.RegisterServer2ServiceServer(s, ss) // 注册服务
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	} else {
		fmt.Printf("server2 started\n")
	}
}
