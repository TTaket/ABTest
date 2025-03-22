package client

import (
	pb "ABTest/pkgs/proto/pb_experiment"
	"context"

	"google.golang.org/grpc"
)

func (c *experimentClient) GetExperiment(ctx context.Context, in *pb.GetExperimentRequest, opts ...grpc.CallOption) (*pb.GetExperimentResponse, error) {
	c.Logger().Infof("GetExperiment client begin: %v", in)

	// logic
	client := pb.NewExperimentServiceClient(c.Conn)
	r, err := client.GetExperiment(ctx, in)
	if err != nil {
		c.Logger().Errorf("%v", err)
	}

	c.Logger().Infof("GetExperiment client end: %v", r)
	return r, err
}
