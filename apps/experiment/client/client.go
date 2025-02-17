package client

import (
	"context"
	"flag"
	"log"

	pb "ABTest/pkgs/proto/pb_experiment"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	addr = flag.String("addr", "localhost:50051", "the address to connect to")
)

// TODO 同一个客户端 多次调用GetExperiment和CreateExperiment 会多次的调用connect
type ExperimentServiceClient interface {
	GetExperiment(ctx context.Context, in *pb.GetExperimentRequest, opts ...grpc.CallOption) (*pb.GetExperimentResponse, error)
	CreateExperiment(ctx context.Context, in *pb.CreateExperimentRequest, opts ...grpc.CallOption) (*pb.CreateExperimentResponse, error)
}

func NewExperimentServiceClient() ExperimentServiceClient {
	return &experimentServiceClient{}
}

type experimentServiceClient struct {
}

func getClientConn() (*grpc.ClientConn, error) {
	conn, err := grpc.NewClient(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	return conn, nil
}
