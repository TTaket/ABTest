package main

import (
	"context"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedABTestServiceServer
}

func (s *server) GetExperiment(ctx context.Context, req *pb.GetExperimentRequest) (*pb.GetExperimentResponse, error) {
	// 实现获取实验的逻辑
	return &pb.GetExperimentResponse{}, nil
}

func (s *server) CreateExperiment(ctx context.Context, req *pb.CreateExperimentRequest) (*pb.CreateExperimentResponse, error) {
	// 实现创建实验的逻辑
	return &pb.CreateExperimentResponse{}, nil
}

func main() {
	// 启动gRPC服务器
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterABTestServiceServer(grpcServer, &server{})
	log.Println("gRPC server listening on port 50051")
	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	// 启动 HTTP 服务器
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	gwMux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err = pb.RegisterABTestServiceHandlerFromEndpoint(ctx, gwMux, "localhost:50051", opts)
	if err != nil {
		log.Fatalf("failed to start HTTP handler: %v", err)
	}

	mux := http.NewServeMux()
	// 挂载静态文件服务
	mux.Handle("/", http.FileServer(http.Dir("./static")))
	// 挂载 gRPC-Gateway 接口到 /api/ 前缀下
	mux.Handle("/api/", http.StripPrefix("/api", gwMux))

	srv := &http.Server{
		Addr:    ":8000",
		Handler: mux,
	}

	go func() {
		log.Println("HTTP server listening on port 8000")
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("HTTP server ListenAndServe: %v", err)
		}
	}()

	// 捕获系统退出信号，实现优雅关停
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	<-stop
	log.Println("Shutting down HTTP server...")

	shutdownCtx, cancelShutdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelShutdown()
	if err := srv.Shutdown(shutdownCtx); err != nil {
		log.Fatalf("HTTP server Shutdown: %v", err)
	}
}
