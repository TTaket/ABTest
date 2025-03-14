package client

import (
	pb "ABTest/pkgs/proto/pb_experiment"
	"context"
	"log"

	"google.golang.org/grpc"
)

func (c *experimentClient) CreateExperiment(ctx context.Context, in *pb.CreateExperimentRequest, opts ...grpc.CallOption) (*pb.CreateExperimentResponse, error) {
	client := pb.NewExperimentServiceClient(c.Conn)
	r, err := client.CreateExperiment(ctx, in)
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	return r, err
}
