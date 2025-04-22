package main

import (
	"context"
	"log"
	"time"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
)

func main() {
	// 1. 创建 gRPC 导出器
	ctx := context.Background()
	exp, err := otlptracegrpc.New(ctx,
		otlptracegrpc.WithInsecure(),                 // 如果使用 TLS，请替换为 WithTLSCredentials
		otlptracegrpc.WithEndpoint("localhost:4317"), // otel-collector 的地址和端口
	)
	if err != nil {
		log.Fatalf("failed to create gRPC exporter: %v", err)
	}

	// 2. 创建资源（Resource）以描述你的服务
	res := resource.NewWithAttributes(
		semconv.SchemaURL,
		semconv.ServiceNameKey.String("my-service"),
		semconv.ServiceVersionKey.String("v0.1.0"),
	)

	// 3. 创建 TracerProvider
	tp := sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(exp), // 使用批处理器
		sdktrace.WithResource(res),
	)

	// 4. 设置全局 TracerProvider
	otel.SetTracerProvider(tp)

	// 5. 创建一个 Tracer 并发送 Span
	tracer := otel.Tracer("my-tracer")
	ctx, span := tracer.Start(ctx, "example-operation")
	defer span.End()

	// 模拟一些工作
	time.Sleep(100 * time.Millisecond)

	// 6. 关闭 TracerProvider 以导出剩余的 Span
	if err := tp.Shutdown(ctx); err != nil {
		log.Fatalf("failed to shutdown TracerProvider: %v", err)
	}
}
