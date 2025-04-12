package service

import (
	"context"

	pb "ABTest/pkgs/proto/pb_%1"
		
	"google.golang.org/grpc"
)

// service 负责业务逻辑实现
type %2Service interface {
	%00002
}

type %3Service struct {

	// others..
}

func New%2Service() %2Service {
	return &%3Service{}
}
