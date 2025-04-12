package client

import (
	pb "ABTest/pkgs/proto/pb_layer"
	"context"

	"google.golang.org/grpc"
)

func (c *layerClient) GetLayer(ctx context.Context, in *pb.GetLayerRequest, opts ...grpc.CallOption) (*pb.GetLayerResponse, error) {
	c.Logger().Infof("GetLayer client begin: %v", in)

	// logic
	client := pb.NewLayerServiceClient(c.Conn)
	r, err := client.GetLayer(ctx, in)
	if err != nil {
		c.Logger().Errorf("%v", err)
	}

	c.Logger().Infof("GetLayer client end: %v", r)
	return r, err
}
