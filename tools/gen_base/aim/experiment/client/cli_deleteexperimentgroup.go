package client

import (
	pb "ABTest/pkgs/proto/pb_experiment"
	"context"

	"google.golang.org/grpc"
)

func (c *experimentClient) DeleteExperimentGroup(ctx context.Context, in *pb.DeleteExperimentGroupRequest, opts ...grpc.CallOption) (*pb.DeleteExperimentGroupResponse, error) {
	c.Logger().Infof("DeleteExperimentGroup client begin: %v", in)

	// logic
	client := pb.NewExperimentServiceClient(c.Conn)
	r, err := client.DeleteExperimentGroup(ctx, in)
	if err != nil {
		c.Logger().Errorf("%v", err)
	}

	c.Logger().Infof("DeleteExperimentGroup client end: %v", r)
	return r, err
}
