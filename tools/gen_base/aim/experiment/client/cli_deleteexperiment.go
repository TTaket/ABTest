package client

import (
	pb "ABTest/pkgs/proto/pb_experiment"
	"context"

	"google.golang.org/grpc"
)

func (c *experimentClient) DeleteExperiment(ctx context.Context, in *pb.DeleteExperimentRequest, opts ...grpc.CallOption) (*pb.DeleteExperimentResponse, error) {
	c.Logger().Infof("DeleteExperiment client begin: %v", in)

	// logic
	client := pb.NewExperimentServiceClient(c.Conn)
	r, err := client.DeleteExperiment(ctx, in)
	if err != nil {
		c.Logger().Errorf("%v", err)
	}

	c.Logger().Infof("DeleteExperiment client end: %v", r)
	return r, err
}
