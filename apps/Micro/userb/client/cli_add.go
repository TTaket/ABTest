package client

import (
	log "ABTest/pkgs/logger"
	pb "ABTest/pkgs/proto/pb_userb"
	"context"
)

func (c *userbClient) AddUser(ctx context.Context, in *pb.AddUserRequest) (*pb.AddUserResponse, error) {
	conn, err := getClientConn()
	if err != nil || conn == nil {
		log.Errorf("did not connect: %v", err)
		return nil, err
	}
	defer conn.Close()
	client := pb.NewUserbServiceClient(conn)
	if client == nil {
		log.Errorf("client is nil")
		return nil, err
	}
	r, err := client.AddUser(ctx, in)
	if err != nil {
		log.Errorf("could not greet: %v", err)
		return nil, err
	}
	return r, err
}

func (c *userbClient) BatchAddUser(ctx context.Context, in *pb.BatchAddUserRequest) (*pb.BatchAddUserResponse, error) {
	conn, err := getClientConn()
	if err != nil || conn == nil {
		log.Errorf("did not connect: %v", err)
		return nil, err
	}
	defer conn.Close()
	client := pb.NewUserbServiceClient(conn)
	r, err := client.BatchAddUser(ctx, in)
	if err != nil {
		log.Errorf("could not greet: %v", err)
		return nil, err
	}
	return r, err
}
