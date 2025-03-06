package client

import (
	pb "ABTest/pkgs/proto/pb_experiment"
	"context"
	"log"

	"google.golang.org/grpc"
)

func (experimentServiceClient) CreateExperiment(ctx context.Context, in *pb.CreateExperimentRequest, opts ...grpc.CallOption) (*pb.CreateExperimentResponse, error) {
	conn, err := getClientConn()
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewExperimentServiceClient(conn)
	r, err := c.CreateExperiment(ctx, in)
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	return r, err
}
