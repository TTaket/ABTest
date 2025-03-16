package client

import (
	"context"

	pb "ABTest/pkgs/proto/pb_userb"
)

func (c *userbClient) DeleteUser(ctx context.Context, in *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error) {
	c.Logger().Infof("DeleteUser client begin: %v", in)

	client := pb.NewUserbServiceClient(c.conn)
	r, err := client.DeleteUser(ctx, in)
	if err != nil {
		c.Logger().Errorf("%v", err)
		return nil, err
	}

	c.Logger().Infof("DeleteUser client end: %v", r)
	return r, err
}

func (c *userbClient) BatchDeleteUser(ctx context.Context, in *pb.BatchDeleteUserRequest) (*pb.BatchDeleteUserResponse, error) {
	c.Logger().Infof("BatchDeleteUser client begin: %v", in)

	client := pb.NewUserbServiceClient(c.conn)
	r, err := client.BatchDeleteUser(ctx, in)
	if err != nil {
		c.Logger().Errorf("%v", err)
		return nil, err
	}

	c.Logger().Infof("BatchDeleteUser client end: %v", r)
	return r, err
}
