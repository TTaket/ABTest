package service

import (
	"context"

	conf "ABTest/apps/Micro/experiment/internal/config"
	pb "ABTest/pkgs/proto/pb_experiment"
)

// implement the interface
func (experimentService) CreateExperiment(ctx context.Context, req *pb.CreateExperimentRequest) (resp *pb.CreateExperimentResponse, err error) {
	conf.Log.Infof("CreateExperiment service begin: %v", req)

	resp = new(pb.CreateExperimentResponse)
	resp.ExperimentId = "1"

	conf.Log.Infof("CreateExperiment service end: %v", resp)
	return resp, nil
}
