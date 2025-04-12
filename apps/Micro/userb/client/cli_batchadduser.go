package client

import (
	pb "ABTest/pkgs/proto/pb_userb"
	"context"

	"google.golang.org/grpc"
)

func (c *userbClient) BatchAddUser(ctx context.Context, in *pb.BatchAddUserRequest, opts ...grpc.CallOption) (*pb.BatchAddUserResponse, error) {
	c.Logger().Infof("BatchAddUser client begin: %v", in)

	// logic
	client := pb.NewUserbServiceClient(c.Conn)
	r, err := client.BatchAddUser(ctx, in)
	if err != nil {
		c.Logger().Errorf("%v", err)
	}

	c.Logger().Infof("BatchAddUser client end: %v", r)
	return r, err
}
