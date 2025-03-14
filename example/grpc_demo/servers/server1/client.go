package server1

import (
	"context"
	tracer "grpcdemo/pkg/tracer"
	pb_server1 "grpcdemo/proto/server1/pb_server1/pb_server1"
	"io"
	"log"
	"time"

	grpc_opentracing "github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"
	"github.com/opentracing/opentracing-go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func setTracer_client1() io.Closer {
	var err error
	var trace opentracing.Tracer
	var iocloser io.Closer
	trace, iocloser, err = tracer.InitJaeger("client1")
	if err != nil {
		log.Fatalf("Could not initialize jaeger tracer: %s", err.Error())
	}
	opentracing.SetGlobalTracer(trace)
	return iocloser
}

type Server1Client interface {
	Get(ctx context.Context, req *pb_server1.GetRequest) (*pb_server1.GetResponse, error)
}

func NewServer1Client() Server1Client {
	return &server1Client{}
}

type server1Client struct {
}

func (cc *server1Client) Get(ctx context.Context, req *pb_server1.GetRequest) (*pb_server1.GetResponse, error) {
	iocloser := setTracer_client1()
	defer iocloser.Close()

	conn, err := grpc.Dial(
		"localhost:50051",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithUnaryInterceptor(
			grpc_opentracing.UnaryClientInterceptor(),
		),
	)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb_server1.NewServer1ServiceClient(conn)

	// Contact the server and print out its response.
	ctxWithTimeout, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()
	r, err := c.Get(ctxWithTimeout, req)
	return r, err
}
