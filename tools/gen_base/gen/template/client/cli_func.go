package client

import (
	pb "ABTest/pkgs/proto/pb_%1"
	"context"

	"google.golang.org/grpc"
)

func (c *experimentClient) %0002(ctx context.Context, in *pb.%0002Request, opts ...grpc.CallOption) (*pb.%0002Response, error) {
	c.Logger().Infof("%0002 client begin: %v", in)

	// logic
	client := pb.%5(c.Conn)
	r, err := client.%0002(ctx, in)
	if err != nil {
		c.Logger().Errorf("%v", err)
	}

	c.Logger().Infof("%0002 client end: %v", r)
	return r, err
}
