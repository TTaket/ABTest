package client

import (
	pb "ABTest/pkgs/proto/pb_dynamic"
	"context"

	"google.golang.org/grpc"
)

func (c *experimentClient) CreateLayerUserBinding(ctx context.Context, in *pb.CreateLayerUserBindingRequest, opts ...grpc.CallOption) (*pb.CreateLayerUserBindingResponse, error) {
	c.Logger().Infof("CreateLayerUserBinding client begin: %v", in)

	// logic
	client := pb.NewDynamicServiceClient(c.Conn)
	r, err := client.CreateLayerUserBinding(ctx, in)
	if err != nil {
		c.Logger().Errorf("%v", err)
	}

	c.Logger().Infof("CreateLayerUserBinding client end: %v", r)
	return r, err
}
