package server2

import (
	"context"
	tracer "grpcdemo/pkg/tracer"
	pb_server2 "grpcdemo/proto/server2/pb_server2/pb_server2"
	"io"
	"log"
	"time"

	grpc_opentracing "github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"
	"github.com/opentracing/opentracing-go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Server2Client interface {
	Get(ctx context.Context, req *pb_server2.GetRequest) (*pb_server2.GetResponse, error)
}

func setTracer_client2() io.Closer {
	var err error
	var trace opentracing.Tracer
	var iocloser io.Closer
	trace, iocloser, err = tracer.InitJaeger("client2")
	if err != nil {
		log.Fatalf("Could not initialize jaeger tracer: %s", err.Error())
	}
	opentracing.SetGlobalTracer(trace)
	return iocloser
}
func NewServer2Client() Server2Client {
	return &server2Client{}
}

type server2Client struct {
}

func (cc *server2Client) Get(ctx context.Context, req *pb_server2.GetRequest) (*pb_server2.GetResponse, error) {
	ioclose := setTracer_client2()
	defer ioclose.Close()

	conn, err := grpc.Dial(
		"localhost:50052",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithUnaryInterceptor(
			grpc_opentracing.UnaryClientInterceptor(),
		),
	)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb_server2.NewServer2ServiceClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()
	r, err := c.Get(ctx, req)
	return r, err
}
