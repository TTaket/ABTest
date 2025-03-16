package service

import (
	"context"

	conf "ABTest/apps/Micro/experiment/internal/config"
	pb "ABTest/pkgs/proto/pb_experiment"
)

// implement the interface
func (experimentService) GetExperiment(ctx context.Context, req *pb.GetExperimentRequest) (resp *pb.GetExperimentResponse, err error) {
	conf.Log.Infof("GetExperiment service begin: %v", req)

	resp = new(pb.GetExperimentResponse)
	resp.ExperimentId = "1"
	resp.Description = "test"
	resp.Name = "mock test"
	conf.Log.Info(req, "\n", resp, "\n")

	conf.Log.Infof("GetExperiment service end: %v", resp)
	return resp, nil
}
