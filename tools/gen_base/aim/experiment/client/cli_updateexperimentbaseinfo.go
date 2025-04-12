package client

import (
	pb "ABTest/pkgs/proto/pb_experiment"
	"context"

	"google.golang.org/grpc"
)

func (c *experimentClient) UpdateExperimentBaseInfo(ctx context.Context, in *pb.UpdateExperimentBaseInfoRequest, opts ...grpc.CallOption) (*pb.UpdateExperimentBaseInfoResponse, error) {
	c.Logger().Infof("UpdateExperimentBaseInfo client begin: %v", in)

	// logic
	client := pb.NewExperimentServiceClient(c.Conn)
	r, err := client.UpdateExperimentBaseInfo(ctx, in)
	if err != nil {
		c.Logger().Errorf("%v", err)
	}

	c.Logger().Infof("UpdateExperimentBaseInfo client end: %v", r)
	return r, err
}
