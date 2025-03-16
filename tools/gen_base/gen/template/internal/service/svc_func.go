package service

import (
	"context"

	conf "ABTest/apps/Micro/%1/internal/config"
	pb "ABTest/pkgs/proto/pb_%1"
)

// implement the interface
func (s *%3Service) %0002(ctx context.Context, req *pb.%0002Request) (resp *pb.%0002Response, err error) {
	conf.Log.Infof("%0002 service begin: %v", req)

	// 这里实现逻辑

	
	conf.Log.Infof("%0002 service end: %v", resp)
	return resp, nil
}
