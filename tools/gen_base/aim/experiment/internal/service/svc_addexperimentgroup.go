package service

import (
	"context"

	conf "ABTest/apps/Micro/experiment/internal/config"
	pb "ABTest/pkgs/proto/pb_experiment"
)

// implement the interface
func (s *experimentService) AddExperimentGroup(ctx context.Context, req *pb.AddExperimentGroupRequest) (resp *pb.AddExperimentGroupResponse, err error) {
	conf.Log.Infof("AddExperimentGroup service begin: %v", req)

	// 这里实现逻辑

	
	conf.Log.Infof("AddExperimentGroup service end: %v", resp)
	return resp, nil
}
