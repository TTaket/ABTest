package xgrpc

import (
	conf "ABTest/pkgs/config"
	"fmt"
	"io"
	"log"

	trace "ABTest/pkgs/tracer"

	"github.com/opentracing/opentracing-go"
	grpc "google.golang.org/grpc"
)

type GrpcClientConfigs struct {
	ServiceName string
	grpc        *conf.Grpc
	Tracing     *Tracing
}
type Tracing struct {
	Tracer  opentracing.Tracer
	Enabled bool `yaml:"enabled"`
}

func NewGrpcClientConfigs(clientname string, grpc *conf.Grpc, jaegercfg *conf.Jaeger) (*GrpcClientConfigs, io.Closer) {
	var (
		tracer opentracing.Tracer
		closer io.Closer
	)
	if jaegercfg.Enabled {
		tracer, closer, _ = trace.NewTracer(clientname, jaegercfg)
	}
	configs := &GrpcClientConfigs{
		grpc:        grpc,
		ServiceName: clientname,
		Tracing:     &Tracing{Tracer: tracer, Enabled: jaegercfg.Enabled},
	}
	return configs, closer
}

func SetClientOpts(c *GrpcClientConfigs) (opts []grpc.DialOption) {
	opts = append(opts, grpc.WithInsecure())
	return opts
}

func GetClientConn(c *GrpcClientConfigs) (*grpc.ClientConn, error) {
	// 配置opt
	opts := SetClientOpts(c)

	conn, err := grpc.Dial(fmt.Sprintf(":%d", c.grpc.Port), opts...)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	return conn, nil
}
