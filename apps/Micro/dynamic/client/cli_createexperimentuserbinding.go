package client

import (
	pb "ABTest/pkgs/proto/pb_dynamic"
	"context"

	"google.golang.org/grpc"
)

func (c *dynamicClient) CreateExperimentUserBinding(ctx context.Context, in *pb.CreateExperimentUserBindingRequest, opts ...grpc.CallOption) (*pb.CreateExperimentUserBindingResponse, error) {
	c.Logger().Infof("CreateExperimentUserBinding client begin: %v", in)

	// logic
	client := pb.NewDynamicServiceClient(c.Conn)
	r, err := client.CreateExperimentUserBinding(ctx, in)
	if err != nil {
		c.Logger().Errorf("%v", err)
	}

	c.Logger().Infof("CreateExperimentUserBinding client end: %v", r)
	return r, err
}
