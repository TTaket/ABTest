package client

import (
	pb "ABTest/pkgs/proto/pb_experiment"
	"context"
	"log"

	"google.golang.org/grpc"
)

func (c *experimentClient) GetExperiment(ctx context.Context, in *pb.GetExperimentRequest, opts ...grpc.CallOption) (*pb.GetExperimentResponse, error) {
	client := pb.NewExperimentServiceClient(c.Conn)
	r, err := client.GetExperiment(ctx, in)
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	return r, err
}
