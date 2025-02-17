package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	pb "ABTest/pkgs/proto/pb_experiment"

	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

// server is used to implement helloworld.GreeterServer.
type server struct {
	pb.UnimplementedExperimentServiceServer
}

// implement the interface
func (server) GetExperiment(ctx context.Context, req *pb.GetExperimentRequest) (resp *pb.GetExperimentResponse, err error) {
	resp = new(pb.GetExperimentResponse)
	resp.ExperimentId = "1"
	resp.Description = "test"
	resp.Name = "mock test"
	return resp, nil
}
func (server) CreateExperiment(ctx context.Context, req *pb.CreateExperimentRequest) (resp *pb.CreateExperimentResponse, err error) {
	resp = new(pb.CreateExperimentResponse)
	resp.ExperimentId = "1"
	return resp, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterExperimentServiceServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
