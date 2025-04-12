package client

import (
	pb "ABTest/pkgs/proto/pb_dynamic"
	"context"

	"google.golang.org/grpc"
)

func (c *experimentClient) GetExperimentLayerBinding(ctx context.Context, in *pb.GetExperimentLayerBindingRequest, opts ...grpc.CallOption) (*pb.GetExperimentLayerBindingResponse, error) {
	c.Logger().Infof("GetExperimentLayerBinding client begin: %v", in)

	// logic
	client := pb.NewDynamicServiceClient(c.Conn)
	r, err := client.GetExperimentLayerBinding(ctx, in)
	if err != nil {
		c.Logger().Errorf("%v", err)
	}

	c.Logger().Infof("GetExperimentLayerBinding client end: %v", r)
	return r, err
}
