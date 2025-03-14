package main

import (
	"context"
	"fmt"
	"io"

	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go/config"
)

func initJaeger(service string) (opentracing.Tracer, io.Closer, error) {
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

func firstFunc(c context.Context, tracer opentracing.Tracer) {
	// 如果tracer为空，初始化一个tracer
	if tracer == nil {
		var err error
		var iocloser io.Closer
		tracer, iocloser, err = initJaeger("firstFunc_trace")
		if err != nil {
			panic(err)
		}
		defer iocloser.Close()
	}

	// 从context中获取span
	spanfa := opentracing.SpanFromContext(c)
	var span opentracing.Span
	if spanfa != nil {
		span = tracer.StartSpan("firstFunc", opentracing.ChildOf(spanfa.Context()))
	} else {
		span = tracer.StartSpan("firstFunc")
	}
	defer span.Finish()
	c = opentracing.ContextWithSpan(c, span)
}

func secondFunc(c context.Context, tracer opentracing.Tracer) {
	// 如果tracer为空，初始化一个tracer
	if tracer == nil {
		var err error
		var iocloser io.Closer
		tracer, iocloser, err = initJaeger("secondFunc_trace")
		if err != nil {
			panic(err)
		}
		defer iocloser.Close()
	}

	// 从context中获取span
	spanfa := opentracing.SpanFromContext(c)
	var span opentracing.Span
	if spanfa != nil {
		span = tracer.StartSpan("secondFunc", opentracing.ChildOf(spanfa.Context()))
	} else {
		span = tracer.StartSpan("secondFunc")
	}
	defer span.Finish()
	c = opentracing.ContextWithSpan(c, span)
}

func firstThenSecond(c context.Context, tracer opentracing.Tracer) {
	// 如果tracer为空，初始化一个tracer
	if tracer == nil {
		var err error
		var iocloser io.Closer
		tracer, iocloser, err = initJaeger("FTS_trace")
		if err != nil {
			panic(err)
		}
		defer iocloser.Close()
	}

	spanfa := opentracing.SpanFromContext(c)
	var span opentracing.Span
	if spanfa != nil {
		span = tracer.StartSpan("firstThenSecond", opentracing.ChildOf(spanfa.Context()))
	} else {
		span = tracer.StartSpan("firstThenSecond")
	}
	defer span.Finish()
	c = opentracing.ContextWithSpan(c, span)
	firstFunc(c, tracer)
	secondFunc(c, tracer)
}

func secondThenFirst(c context.Context, tracer opentracing.Tracer) {
	// 如果tracer为空，初始化一个tracer
	if tracer == nil {
		var err error
		var iocloser io.Closer
		tracer, iocloser, err = initJaeger("STF_trace")
		if err != nil {
			panic(err)
		}
		defer iocloser.Close()
	}

	spanfa := opentracing.SpanFromContext(c)
	var span opentracing.Span
	if spanfa != nil {
		span = tracer.StartSpan("secondThenFirst", opentracing.ChildOf(spanfa.Context()))
	} else {
		span = tracer.StartSpan("secondThenFirst")
	}
	defer span.Finish()
	c = opentracing.ContextWithSpan(c, span)
	secondFunc(c, tracer)
	firstFunc(c, tracer)
}
func main() {
	ctx := context.Background()
	for {
		var input int
		_, err := fmt.Scan(&input)
		if err != nil {
			fmt.Println("Error reading input:", err)
			continue
		}

		switch input {
		case 1:
			firstFunc(ctx, nil)
		case 2:
			secondFunc(ctx, nil)
		case 3:
			firstThenSecond(ctx, nil)
		case 4:
			secondThenFirst(ctx, nil)
		default:
			fmt.Println("Invalid input, please enter a number between 1 and 4")
		}
	}
}
