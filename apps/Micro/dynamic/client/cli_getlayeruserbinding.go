package client

import (
	pb "ABTest/pkgs/proto/pb_dynamic"
	"context"

	"google.golang.org/grpc"
)

func (c *dynamicClient) GetLayerUserBinding(ctx context.Context, in *pb.GetLayerUserBindingRequest, opts ...grpc.CallOption) (*pb.GetLayerUserBindingResponse, error) {
	c.Logger().Infof("GetLayerUserBinding client begin: %v", in)

	// logic
	client := pb.NewDynamicServiceClient(c.Conn)
	r, err := client.GetLayerUserBinding(ctx, in)
	if err != nil {
		c.Logger().Errorf("%v", err)
	}

	c.Logger().Infof("GetLayerUserBinding client end: %v", r)
	return r, err
}
