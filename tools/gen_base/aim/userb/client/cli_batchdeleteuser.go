package client

import (
	pb "ABTest/pkgs/proto/pb_userb"
	"context"

	"google.golang.org/grpc"
)

func (c *experimentClient) BatchDeleteUser(ctx context.Context, in *pb.BatchDeleteUserRequest, opts ...grpc.CallOption) (*pb.BatchDeleteUserResponse, error) {
	c.Logger().Infof("BatchDeleteUser client begin: %v", in)

	// logic
	client := pb.NewUserbServiceClient(c.Conn)
	r, err := client.BatchDeleteUser(ctx, in)
	if err != nil {
		c.Logger().Errorf("%v", err)
	}

	c.Logger().Infof("BatchDeleteUser client end: %v", r)
	return r, err
}
