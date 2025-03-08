package client

import (
	"context"

	log "ABTest/pkgs/logger"
	pb "ABTest/pkgs/proto/pb_userb"
)

func (c *userbClient) DeleteUser(ctx context.Context, in *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error) {
	conn, err := getClientConn()
	if err != nil || conn == nil {
		log.Errorf("did not connect: %v", err)
		return nil, err
	}
	defer conn.Close()
	client := pb.NewUserbServiceClient(conn)
	r, err := client.DeleteUser(ctx, in)
	if err != nil {
		log.Errorf("could not greet: %v", err)
		return nil, err
	}
	return r, err
}

func (c *userbClient) BatchDeleteUser(ctx context.Context, in *pb.BatchDeleteUserRequest) (*pb.BatchDeleteUserResponse, error) {
	conn, err := getClientConn()
	if err != nil || conn == nil {
		log.Errorf("did not connect: %v", err)
		return nil, err
	}
	defer conn.Close()
	client := pb.NewUserbServiceClient(conn)
	r, err := client.BatchDeleteUser(ctx, in)
	if err != nil {
		log.Errorf("could not greet: %v", err)
		return nil, err
	}
	return r, err
}
