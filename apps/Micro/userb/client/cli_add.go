package client

import (
	pb "ABTest/pkgs/proto/pb_userb"
	"context"
)

func (c *userbClient) AddUser(ctx context.Context, in *pb.AddUserRequest) (*pb.AddUserResponse, error) {
	c.Logger().Infof("AddUser client begin: %v", in)

	client := pb.NewUserbServiceClient(c.conn)
	r, err := client.AddUser(ctx, in)
	if err != nil {
		c.Logger().Errorf("%v", err)
		return nil, err
	}

	c.Logger().Infof("AddUser client end: %v", r)
	return r, err
}

func (c *userbClient) BatchAddUser(ctx context.Context, in *pb.BatchAddUserRequest) (*pb.BatchAddUserResponse, error) {
	c.Logger().Infof("BatchAddUser client begin: %v", in)

	client := pb.NewUserbServiceClient(c.conn)
	r, err := client.BatchAddUser(ctx, in)
	if err != nil {
		c.Logger().Errorf("%v", err)
		return nil, err
	}

	c.Logger().Infof("BatchAddUser client end: %v", r)
	return r, err
}
