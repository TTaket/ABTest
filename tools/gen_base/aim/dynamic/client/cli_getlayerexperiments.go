package client

import (
	pb "ABTest/pkgs/proto/pb_dynamic"
	"context"

	"google.golang.org/grpc"
)

func (c *experimentClient) GetLayerExperiments(ctx context.Context, in *pb.GetLayerExperimentsRequest, opts ...grpc.CallOption) (*pb.GetLayerExperimentsResponse, error) {
	c.Logger().Infof("GetLayerExperiments client begin: %v", in)

	// logic
	client := pb.NewDynamicServiceClient(c.Conn)
	r, err := client.GetLayerExperiments(ctx, in)
	if err != nil {
		c.Logger().Errorf("%v", err)
	}

	c.Logger().Infof("GetLayerExperiments client end: %v", r)
	return r, err
}
