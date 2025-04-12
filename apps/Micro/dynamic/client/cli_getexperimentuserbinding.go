package client

import (
	pb "ABTest/pkgs/proto/pb_dynamic"
	"context"

	"google.golang.org/grpc"
)

func (c *dynamicClient) GetExperimentUserBinding(ctx context.Context, in *pb.GetExperimentUserBindingRequest, opts ...grpc.CallOption) (*pb.GetExperimentUserBindingResponse, error) {
	c.Logger().Infof("GetExperimentUserBinding client begin: %v", in)

	// logic
	client := pb.NewDynamicServiceClient(c.Conn)
	r, err := client.GetExperimentUserBinding(ctx, in)
	if err != nil {
		c.Logger().Errorf("%v", err)
	}

	c.Logger().Infof("GetExperimentUserBinding client end: %v", r)
	return r, err
}
