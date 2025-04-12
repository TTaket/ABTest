package client

import (
	pb "ABTest/pkgs/proto/pb_userb"
	"context"

	"google.golang.org/grpc"
)

func (c *experimentClient) ScatterTraffic(ctx context.Context, in *pb.ScatterTrafficRequest, opts ...grpc.CallOption) (*pb.ScatterTrafficResponse, error) {
	c.Logger().Infof("ScatterTraffic client begin: %v", in)

	// logic
	client := pb.NewUserbServiceClient(c.Conn)
	r, err := client.ScatterTraffic(ctx, in)
	if err != nil {
		c.Logger().Errorf("%v", err)
	}

	c.Logger().Infof("ScatterTraffic client end: %v", r)
	return r, err
}
