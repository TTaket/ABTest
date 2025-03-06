package xgrpc

import (
	conf "ABTest/pkgs/config"
	log "ABTest/pkgs/logger"
	trace "ABTest/pkgs/tracer"
	"ABTest/pkgs/xetcd"
	"fmt"
	"io"
	"time"

	otgrpc "github.com/opentracing-contrib/go-grpc"
	"github.com/opentracing/opentracing-go"
	grpc "google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/resolver"
)

type GrpcClientConfigs struct {
	ServiceName string
	grpc        *conf.Grpc
	Tracing     *Tracing
	Etcd        *conf.Etcd
}
type Tracing struct {
	Tracer  opentracing.Tracer
	Enabled bool `yaml:"enabled"`
}

func NewGrpcClientConfigs(clientname string, grpc *conf.Grpc, etcd *conf.Etcd, jaegercfg *conf.Jaeger) (*GrpcClientConfigs, io.Closer) {
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
		Etcd:        etcd,
	}
	return configs, closer
}

func SetClientOpts(c *GrpcClientConfigs) (opts []grpc.DialOption) {
	opts = append(opts, grpc.WithInsecure())
	// tracing
	if c.Tracing.Enabled {
		opts = append(opts, grpc.WithUnaryInterceptor(otgrpc.OpenTracingClientInterceptor(c.Tracing.Tracer)))
	}
	// keepalive
	keepParams := grpc.WithKeepaliveParams(keepalive.ClientParameters{
		Time:    time.Duration(c.grpc.Time),
		Timeout: time.Duration(c.grpc.Timeout),
	})
	opts = append(opts, keepParams)
	// etcd
	if c.Etcd != nil {
		// etcd 连接
		etcdClient, err := xetcd.GetEtcdClient(c.Etcd)
		if err != nil {
			log.Errorf("etcd client error: %v", err)
		}
		resolver.Register(xetcd.NewEtcdBuilder(etcdClient, c.Etcd.Schema, c.ServiceName))
		// etcd 轮询策略
		serviceConfig := fmt.Sprintf("{\"loadBalancingPolicy\": \"%s\"}", c.Etcd.BalancingPolicy)
		opts = append(opts, grpc.WithDefaultServiceConfig(serviceConfig))
	}
	return opts
}

func GetClientConn(c *GrpcClientConfigs) (*grpc.ClientConn, error) {
	// 配置opt
	opts := SetClientOpts(c)

	// etcd 连接
	if c.Etcd != nil {
		conn, err := grpc.Dial(fmt.Sprintf("%s:///%s", c.Etcd.Schema, c.ServiceName), opts...)
		if err != nil {
			log.Errorf("did not connect by etcd : %v", err)
		}
		return conn, nil

	} else {
		// 连接
		conn, err := grpc.Dial(fmt.Sprintf(":%d", c.grpc.Port), opts...)
		if err != nil {
			log.Errorf("did not connect: %v", err)
		}
		return conn, nil
	}
}
