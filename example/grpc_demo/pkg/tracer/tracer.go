package tracer

import (
	"io"

	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go/config"
)

func InitJaeger(service string) (opentracing.Tracer, io.Closer, error) {
	// 创建一个Jaeger配置
	cfg := config.Configuration{
		ServiceName: service,
		Sampler: &config.SamplerConfig{
			Type:  "const",
			Param: 1,
		},
		Reporter: &config.ReporterConfig{
			LogSpans:           true,
			LocalAgentHostPort: "127.0.0.1:6831",
		},
	}
	// 初始化配置
	tracer, closer, err := cfg.NewTracer()
	if err != nil {
		panic(err)
	}
	return tracer, closer, err
}
