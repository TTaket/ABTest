package client

import (
	pb "ABTest/pkgs/proto/pb_layer"
	"context"

	"google.golang.org/grpc"
)

func (c *experimentClient) DeleteLayer(ctx context.Context, in *pb.DeleteLayerRequest, opts ...grpc.CallOption) (*pb.DeleteLayerResponse, error) {
	c.Logger().Infof("DeleteLayer client begin: %v", in)

	// logic
	client := pb.NewLayerServiceClient(c.Conn)
	r, err := client.DeleteLayer(ctx, in)
	if err != nil {
		c.Logger().Errorf("%v", err)
	}

	c.Logger().Infof("DeleteLayer client end: %v", r)
	return r, err
}
