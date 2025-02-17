package server

import (
	"ABTest/apps/experiment/internal/service"
	pb "ABTest/pkgs/proto/pb_experiment"
	"flag"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

// server 负责完成逻辑和grpc的交互
type ExperimentServer interface {
	Run()
}

// NewExperimentServer returns a new ExperimentServer
func NewExperimentServer() ExperimentServer {
	return &experimentServer{
		experimentService: service.NewExperimentService(),
	}
}

type experimentServer struct {
	pb.UnimplementedExperimentServiceServer
	experimentService service.ExperimentService
}

// Run starts the gRPC server
func (s *experimentServer) Run() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterExperimentServiceServer(grpcServer, s)
	log.Printf("server listening at %v", lis.Addr())
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
