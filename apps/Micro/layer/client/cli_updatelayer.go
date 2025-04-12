package client

import (
	pb "ABTest/pkgs/proto/pb_layer"
	"context"

	"google.golang.org/grpc"
)

func (c *layerClient) UpdateLayer(ctx context.Context, in *pb.UpdateLayerRequest, opts ...grpc.CallOption) (*pb.UpdateLayerResponse, error) {
	c.Logger().Infof("UpdateLayer client begin: %v", in)

	// logic
	client := pb.NewLayerServiceClient(c.Conn)
	r, err := client.UpdateLayer(ctx, in)
	if err != nil {
		c.Logger().Errorf("%v", err)
	}

	c.Logger().Infof("UpdateLayer client end: %v", r)
	return r, err
}
