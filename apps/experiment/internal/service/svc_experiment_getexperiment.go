package service

import (
	"context"

	pb "ABTest/pkgs/proto/pb_experiment"
)

// implement the interface
func (experimentService) GetExperiment(ctx context.Context, req *pb.GetExperimentRequest) (resp *pb.GetExperimentResponse, err error) {
	resp = new(pb.GetExperimentResponse)
	resp.ExperimentId = "1"
	resp.Description = "test"
	resp.Name = "mock test"
	return resp, nil
}
