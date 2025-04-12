package client

import (
	pb "ABTest/pkgs/proto/pb_dynamic"
	"context"

	"google.golang.org/grpc"
)

func (c *dynamicClient) CreateExperimentLayerBinding(ctx context.Context, in *pb.CreateExperimentLayerBindingRequest, opts ...grpc.CallOption) (*pb.CreateExperimentLayerBindingResponse, error) {
	c.Logger().Infof("CreateExperimentLayerBinding client begin: %v", in)

	// logic
	client := pb.NewDynamicServiceClient(c.Conn)
	r, err := client.CreateExperimentLayerBinding(ctx, in)
	if err != nil {
		c.Logger().Errorf("%v", err)
	}

	c.Logger().Infof("CreateExperimentLayerBinding client end: %v", r)
	return r, err
}
