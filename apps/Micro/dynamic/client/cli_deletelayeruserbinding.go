package client

import (
	pb "ABTest/pkgs/proto/pb_dynamic"
	"context"

	"google.golang.org/grpc"
)

func (c *dynamicClient) DeleteLayerUserBinding(ctx context.Context, in *pb.DeleteLayerUserBindingRequest, opts ...grpc.CallOption) (*pb.DeleteLayerUserBindingResponse, error) {
	c.Logger().Infof("DeleteLayerUserBinding client begin: %v", in)

	// logic
	client := pb.NewDynamicServiceClient(c.Conn)
	r, err := client.DeleteLayerUserBinding(ctx, in)
	if err != nil {
		c.Logger().Errorf("%v", err)
	}

	c.Logger().Infof("DeleteLayerUserBinding client end: %v", r)
	return r, err
}
