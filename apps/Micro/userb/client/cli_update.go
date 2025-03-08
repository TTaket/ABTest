package client

import (
	"context"

	log "ABTest/pkgs/logger"
	pb "ABTest/pkgs/proto/pb_userb"
)

func (c *userbClient) UpdateUser(ctx context.Context, in *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error) {
	conn, err := getClientConn()
	if err != nil || conn == nil {
		log.Errorf("did not connect: %v", err)
		return nil, err
	}
	defer conn.Close()
	client := pb.NewUserbServiceClient(conn)
	r, err := client.UpdateUser(ctx, in)
	if err != nil {
		log.Errorf("could not greet: %v", err)
		return nil, err
	}
	return r, err
}

func (c *userbClient) BatchUpdateUser(ctx context.Context, in *pb.BatchUpdateUserRequest) (*pb.BatchUpdateUserResponse, error) {
	conn, err := getClientConn()
	if err != nil || conn == nil {
		log.Errorf("did not connect: %v", err)
		return nil, err
	}
	defer conn.Close()
	client := pb.NewUserbServiceClient(conn)
	r, err := client.BatchUpdateUser(ctx, in)
	if err != nil {
		log.Errorf("could not greet: %v", err)
		return nil, err
	}
	return r, err
}
