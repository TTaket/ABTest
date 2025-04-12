package client

import (
	pb "ABTest/pkgs/proto/pb_userb"
	"context"

	"google.golang.org/grpc"
)

func (c *userbClient) GetUserInfo(ctx context.Context, in *pb.GetUserInfoRequest, opts ...grpc.CallOption) (*pb.GetUserInfoResponse, error) {
	c.Logger().Infof("GetUserInfo client begin: %v", in)

	// logic
	client := pb.NewUserbServiceClient(c.Conn)
	r, err := client.GetUserInfo(ctx, in)
	if err != nil {
		c.Logger().Errorf("%v", err)
	}

	c.Logger().Infof("GetUserInfo client end: %v", r)
	return r, err
}
