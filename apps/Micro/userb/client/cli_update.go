package client

import (
	"context"

	pb "ABTest/pkgs/proto/pb_userb"
)

func (c *userbClient) UpdateUser(ctx context.Context, in *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error) {
	c.Logger().Infof("UpdateUser client begin: %v", in)

	client := pb.NewUserbServiceClient(c.conn)
	r, err := client.UpdateUser(ctx, in)
	if err != nil {
		c.Logger().Errorf("%v", err)
		return nil, err
	}

	c.Logger().Infof("UpdateUser client end: %v", r)
	return r, err
}

func (c *userbClient) BatchUpdateUser(ctx context.Context, in *pb.BatchUpdateUserRequest) (*pb.BatchUpdateUserResponse, error) {
	c.Logger().Infof("BatchUpdateUser client begin: %v", in)

	client := pb.NewUserbServiceClient(c.conn)
	r, err := client.BatchUpdateUser(ctx, in)
	if err != nil {
		c.Logger().Errorf("%v", err)
		return nil, err
	}

	c.Logger().Infof("BatchUpdateUser client end: %v", r)
	return r, err
}
