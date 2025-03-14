package service

import (
	"context"
	"fmt"

	userbClient "ABTest/apps/Micro/userb/client"
	log "ABTest/pkgs/logger"
	pb "ABTest/pkgs/proto/pb_experiment"
	"ABTest/pkgs/proto/pb_userb"

	"google.golang.org/protobuf/proto"
)

// implement the interface
func (experimentService) GetExperiment(ctx context.Context, req *pb.GetExperimentRequest) (resp *pb.GetExperimentResponse, err error) {
	resp = new(pb.GetExperimentResponse)
	resp.ExperimentId = "1"
	resp.Description = "test"
	resp.Name = "mock test"
	log.Info(req, "\n", resp, "\n")

	//Todo: temp
	{
		UserbClient := userbClient.NewUserbClient()
		resp, err := UserbClient.GetUserInfo(ctx, &pb_userb.GetUserInfoRequest{
			UserInfo: &pb_userb.UserInfo{
				UserId: proto.Uint64(1),
			},
		})
		if err != nil {
			panic(err)
		}
		fmt.Printf("TestDemo: get user info %s\n", resp.String())
	}

	return resp, nil
}
