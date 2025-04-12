package client

import (
	pb "ABTest/pkgs/proto/pb_userb"
	"context"

	"google.golang.org/grpc"
)

func (c *experimentClient) BatchGetUserInfo(ctx context.Context, in *pb.BatchGetUserInfoRequest, opts ...grpc.CallOption) (*pb.BatchGetUserInfoResponse, error) {
	c.Logger().Infof("BatchGetUserInfo client begin: %v", in)

	// logic
	client := pb.NewUserbServiceClient(c.Conn)
	r, err := client.BatchGetUserInfo(ctx, in)
	if err != nil {
		c.Logger().Errorf("%v", err)
	}

	c.Logger().Infof("BatchGetUserInfo client end: %v", r)
	return r, err
}
