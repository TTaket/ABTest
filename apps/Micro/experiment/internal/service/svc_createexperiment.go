package service

import (
	"context"

	pb "ABTest/pkgs/proto/pb_experiment"
)

// implement the interface
func (experimentService) CreateExperiment(ctx context.Context, req *pb.CreateExperimentRequest) (resp *pb.CreateExperimentResponse, err error) {
	resp = new(pb.CreateExperimentResponse)
	resp.ExperimentId = "1"
	return resp, nil
}
