package client

import (
	log "ABTest/pkgs/logger"
	pb "ABTest/pkgs/proto/pb_userb"
	"context"
)

func (c *userbClient) AddUser(ctx context.Context, in *pb.AddUserRequest) (*pb.AddUserResponse, error) {
	client := pb.NewUserbServiceClient(c.conn)
	r, err := client.AddUser(ctx, in)
	if err != nil {
		log.Errorf("could not greet: %v", err)
		return nil, err
	}
	return r, err
}

func (c *userbClient) BatchAddUser(ctx context.Context, in *pb.BatchAddUserRequest) (*pb.BatchAddUserResponse, error) {
	client := pb.NewUserbServiceClient(c.conn)
	r, err := client.BatchAddUser(ctx, in)
	if err != nil {
		log.Errorf("could not greet: %v", err)
		return nil, err
	}
	return r, err
}
