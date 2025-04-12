package client

import (
	pb "ABTest/pkgs/proto/pb_experiment"
	"context"

	"google.golang.org/grpc"
)

func (c *experimentClient) CreateExperiment(ctx context.Context, in *pb.CreateExperimentRequest, opts ...grpc.CallOption) (*pb.CreateExperimentResponse, error) {
	c.Logger().Infof("CreateExperiment client begin: %v", in)

	// logic
	client := pb.NewExperimentServiceClient(c.Conn)
	r, err := client.CreateExperiment(ctx, in)
	if err != nil {
		c.Logger().Errorf("%v", err)
	}

	c.Logger().Infof("CreateExperiment client end: %v", r)
	return r, err
}
