package server1

import (
	"context"
	"fmt"
	tracer "grpcdemo/pkg/tracer"
	pb_server1 "grpcdemo/proto/server1/pb_server1/pb_server1"
	"grpcdemo/proto/server2/pb_server2/pb_server2"
	"grpcdemo/servers/server2"
	"io"
	"log"
	"net"

	grpc_opentracing "github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"
	"github.com/opentracing/opentracing-go"
	"google.golang.org/grpc"
)

// 直接实现 pb_server1.Server1ServiceServer 接口
type server1Service struct {
	pb_server1.UnimplementedServer1ServiceServer // 嵌入 UnimplementedServer1ServiceServer
}

func setTracer_server1() io.Closer {
	trace, iocloser, err := tracer.InitJaeger("server1")
	if err != nil {
		log.Fatalf("Could not initialize jaeger tracer: %s", err.Error())
	}
	opentracing.SetGlobalTracer(trace)
	return iocloser
}
func (s *server1Service) Get(ctx context.Context, req *pb_server1.GetRequest) (*pb_server1.GetResponse, error) {

	resp := &pb_server1.GetResponse{
		ExperimentId: req.GetExperimentId(),
	}

	// Call server2
	{
		r, err := server2.NewServer2Client().Get(ctx, &pb_server2.GetRequest{
			ExperimentId: "1",
		})
		if err != nil {
			panic(err)
		}
		fmt.Printf("server1 call server2 %s \n", r.String())
	}

	return resp, nil
}

func StartServer() {
	ioclose := setTracer_server1()
	defer ioclose.Close()

	// Start the server
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer(
		grpc.UnaryInterceptor(
			grpc_opentracing.UnaryServerInterceptor(),
		),
	)
	ss := &server1Service{}                        // 直接创建 server1Service 实例
	pb_server1.RegisterServer1ServiceServer(s, ss) // 注册服务
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	} else {
		fmt.Printf("server1 started\n")
	}
}
