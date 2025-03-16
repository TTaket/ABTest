package client

import (
	"context"

	pb "ABTest/pkgs/proto/pb_userb"
)

func (c *userbClient) GetUserInfo(ctx context.Context, in *pb.GetUserInfoRequest) (*pb.GetUserInfoResponse, error) {
	c.Logger().Infof("GetUserInfo client begin: %v", in)

	client := pb.NewUserbServiceClient(c.conn)
	r, err := client.GetUserInfo(ctx, in)
	if err != nil {
		c.Logger().Errorf("%v", err)
		return nil, err
	}

	c.Logger().Infof("GetUserInfo client end: %v", r)
	return r, err
}

func (c *userbClient) BatchGetUserInfo(ctx context.Context, in *pb.BatchGetUserInfoRequest) (*pb.BatchGetUserInfoResponse, error) {
	c.Logger().Infof("BatchGetUserInfo client begin: %v", in)

	client := pb.NewUserbServiceClient(c.conn)
	r, err := client.BatchGetUserInfo(ctx, in)
	if err != nil {
		c.Logger().Errorf("%v", err)
		return nil, err
	}

	c.Logger().Infof("BatchGetUserInfo client end: %v", r)
	return r, err
}
