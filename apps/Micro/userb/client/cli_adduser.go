package client

import (
	pb "ABTest/pkgs/proto/pb_userb"
	"context"

	"google.golang.org/grpc"
)

func (c *userbClient) AddUser(ctx context.Context, in *pb.AddUserRequest, opts ...grpc.CallOption) (*pb.AddUserResponse, error) {
	c.Logger().Infof("AddUser client begin: %v", in)

	// logic
	client := pb.NewUserbServiceClient(c.Conn)
	r, err := client.AddUser(ctx, in)
	if err != nil {
		c.Logger().Errorf("%v", err)
	}

	c.Logger().Infof("AddUser client end: %v", r)
	return r, err
}
