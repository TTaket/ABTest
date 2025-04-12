package service

import (
	"context"

	conf "ABTest/apps/Micro/experiment/internal/config"
	pb "ABTest/pkgs/proto/pb_experiment"
)

// implement the interface
func (s *experimentService) DeleteExperimentGroup(ctx context.Context, req *pb.DeleteExperimentGroupRequest) (resp *pb.DeleteExperimentGroupResponse, err error) {
	conf.Log.Infof("DeleteExperimentGroup service begin: %v", req)

	// 这里实现逻辑

	
	conf.Log.Infof("DeleteExperimentGroup service end: %v", resp)
	return resp, nil
}
