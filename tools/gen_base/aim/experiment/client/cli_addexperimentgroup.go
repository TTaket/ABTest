package client

import (
	pb "ABTest/pkgs/proto/pb_experiment"
	"context"

	"google.golang.org/grpc"
)

func (c *experimentClient) AddExperimentGroup(ctx context.Context, in *pb.AddExperimentGroupRequest, opts ...grpc.CallOption) (*pb.AddExperimentGroupResponse, error) {
	c.Logger().Infof("AddExperimentGroup client begin: %v", in)

	// logic
	client := pb.NewExperimentServiceClient(c.Conn)
	r, err := client.AddExperimentGroup(ctx, in)
	if err != nil {
		c.Logger().Errorf("%v", err)
	}

	c.Logger().Infof("AddExperimentGroup client end: %v", r)
	return r, err
}
