package client

import (
	pb "ABTest/pkgs/proto/pb_dynamic"
	"context"

	"google.golang.org/grpc"
)

func (c *dynamicClient) DeleteExperimentUserBinding(ctx context.Context, in *pb.DeleteExperimentUserBindingRequest, opts ...grpc.CallOption) (*pb.DeleteExperimentUserBindingResponse, error) {
	c.Logger().Infof("DeleteExperimentUserBinding client begin: %v", in)

	// logic
	client := pb.NewDynamicServiceClient(c.Conn)
	r, err := client.DeleteExperimentUserBinding(ctx, in)
	if err != nil {
		c.Logger().Errorf("%v", err)
	}

	c.Logger().Infof("DeleteExperimentUserBinding client end: %v", r)
	return r, err
}
