package client

import (
	pb "ABTest/pkgs/proto/pb_userb"
	"context"

	"google.golang.org/grpc"
)

func (c *experimentClient) BatchUpdateUser(ctx context.Context, in *pb.BatchUpdateUserRequest, opts ...grpc.CallOption) (*pb.BatchUpdateUserResponse, error) {
	c.Logger().Infof("BatchUpdateUser client begin: %v", in)

	// logic
	client := pb.NewUserbServiceClient(c.Conn)
	r, err := client.BatchUpdateUser(ctx, in)
	if err != nil {
		c.Logger().Errorf("%v", err)
	}

	c.Logger().Infof("BatchUpdateUser client end: %v", r)
	return r, err
}
