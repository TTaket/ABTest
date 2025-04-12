package client

import (
	pb "ABTest/pkgs/proto/pb_layer"
	"context"

	"google.golang.org/grpc"
)

func (c *experimentClient) ListLayers(ctx context.Context, in *pb.ListLayersRequest, opts ...grpc.CallOption) (*pb.ListLayersResponse, error) {
	c.Logger().Infof("ListLayers client begin: %v", in)

	// logic
	client := pb.NewLayerServiceClient(c.Conn)
	r, err := client.ListLayers(ctx, in)
	if err != nil {
		c.Logger().Errorf("%v", err)
	}

	c.Logger().Infof("ListLayers client end: %v", r)
	return r, err
}
