package client

import (
	pb "ABTest/pkgs/proto/pb_layer"
	"context"

	"google.golang.org/grpc"
)

func (c *experimentClient) CreateLayer(ctx context.Context, in *pb.CreateLayerRequest, opts ...grpc.CallOption) (*pb.CreateLayerResponse, error) {
	c.Logger().Infof("CreateLayer client begin: %v", in)

	// logic
	client := pb.NewLayerServiceClient(c.Conn)
	r, err := client.CreateLayer(ctx, in)
	if err != nil {
		c.Logger().Errorf("%v", err)
	}

	c.Logger().Infof("CreateLayer client end: %v", r)
	return r, err
}
