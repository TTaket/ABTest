package client

import (
	"context"

	log "ABTest/pkgs/logger"
	pb "ABTest/pkgs/proto/pb_userb"
)

func (c *userbClient) GetUserInfo(ctx context.Context, in *pb.GetUserInfoRequest) (*pb.GetUserInfoResponse, error) {
	client := pb.NewUserbServiceClient(c.conn)
	r, err := client.GetUserInfo(ctx, in)
	if err != nil {
		log.Errorf("could not greet: %v", err)
		return nil, err
	}
	return r, err
}

func (c *userbClient) BatchGetUserInfo(ctx context.Context, in *pb.BatchGetUserInfoRequest) (*pb.BatchGetUserInfoResponse, error) {
	client := pb.NewUserbServiceClient(c.conn)
	r, err := client.BatchGetUserInfo(ctx, in)
	if err != nil {
		log.Errorf("could not greet: %v", err)
		return nil, err
	}
	return r, err
}
