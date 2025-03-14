package client

import (
	"context"

	log "ABTest/pkgs/logger"
	pb "ABTest/pkgs/proto/pb_userb"
)

func (c *userbClient) DeleteUser(ctx context.Context, in *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error) {
	client := pb.NewUserbServiceClient(c.conn)
	r, err := client.DeleteUser(ctx, in)
	if err != nil {
		log.Errorf("could not greet: %v", err)
		return nil, err
	}
	return r, err
}

func (c *userbClient) BatchDeleteUser(ctx context.Context, in *pb.BatchDeleteUserRequest) (*pb.BatchDeleteUserResponse, error) {
	client := pb.NewUserbServiceClient(c.conn)
	r, err := client.BatchDeleteUser(ctx, in)
	if err != nil {
		log.Errorf("could not greet: %v", err)
		return nil, err
	}
	return r, err
}
