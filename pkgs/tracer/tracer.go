package tracer

import (
	conf "ABTest/pkgs/config"
	"fmt"
	"io"
	"log"
	"os"

	opentracing "github.com/opentracing/opentracing-go"
	jaegercfg "github.com/uber/jaeger-client-go/config"
)

func NewTracer(servicename string, c *conf.Jaeger) (tracer opentracing.Tracer, closer io.Closer, err error) {
	// 1. 打开日志文件
	logDir := fmt.Sprintf("%s%s", c.LogPath, servicename+"/")
	if err := os.MkdirAll(logDir, 0755); err != nil {
		return nil, nil, fmt.Errorf("failed to create log directory: %w", err)
	}
	logFile, err := os.OpenFile(fmt.Sprintf("%sjaeger.log", logDir), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to open log file: %w", err)
	}

	// 2. 创建自定义 Logger
	jaegerLogger := &customLogger{
		logger: log.New(io.MultiWriter(logFile), // 输出到文件
			"",                                  // 前缀
			log.Ldate|log.Ltime|log.Lshortfile), // 日志格式
	}
	// TODO 自动清理日志文件
	cfg := &jaegercfg.Configuration{
		ServiceName: servicename,
		Sampler: &jaegercfg.SamplerConfig{
			Type:  c.SamplerType,
			Param: c.Param,
		},
		Reporter: &jaegercfg.ReporterConfig{
			LogSpans:           c.LogSpans,
			LocalAgentHostPort: c.HostPort,
		},
	}

	tracer, closer, err = cfg.NewTracer(jaegercfg.Logger(jaegerLogger))
	return
}

// 自定义 Logger 实现
type customLogger struct {
	logger *log.Logger
}

func (l *customLogger) Error(msg string) {
	l.logger.Printf("ERROR: %s", msg)
}

func (l *customLogger) Infof(msg string, args ...interface{}) {
	l.logger.Printf("INFO: "+msg, args...)
}

func (l *customLogger) Debugf(msg string, args ...interface{}) {
	l.logger.Printf("DEBUG: "+msg, args...)
}
