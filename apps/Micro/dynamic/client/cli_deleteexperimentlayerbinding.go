package client

import (
	pb "ABTest/pkgs/proto/pb_dynamic"
	"context"

	"google.golang.org/grpc"
)

func (c *dynamicClient) DeleteExperimentLayerBinding(ctx context.Context, in *pb.DeleteExperimentLayerBindingRequest, opts ...grpc.CallOption) (*pb.DeleteExperimentLayerBindingResponse, error) {
	c.Logger().Infof("DeleteExperimentLayerBinding client begin: %v", in)

	// logic
	client := pb.NewDynamicServiceClient(c.Conn)
	r, err := client.DeleteExperimentLayerBinding(ctx, in)
	if err != nil {
		c.Logger().Errorf("%v", err)
	}

	c.Logger().Infof("DeleteExperimentLayerBinding client end: %v", r)
	return r, err
}
