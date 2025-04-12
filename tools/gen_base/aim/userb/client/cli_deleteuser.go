package client

import (
	pb "ABTest/pkgs/proto/pb_userb"
	"context"

	"google.golang.org/grpc"
)

func (c *experimentClient) DeleteUser(ctx context.Context, in *pb.DeleteUserRequest, opts ...grpc.CallOption) (*pb.DeleteUserResponse, error) {
	c.Logger().Infof("DeleteUser client begin: %v", in)

	// logic
	client := pb.NewUserbServiceClient(c.Conn)
	r, err := client.DeleteUser(ctx, in)
	if err != nil {
		c.Logger().Errorf("%v", err)
	}

	c.Logger().Infof("DeleteUser client end: %v", r)
	return r, err
}
