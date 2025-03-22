package main

import (
	client "ABTest/apps/Micro/experiment/client"
	pb "ABTest/pkgs/proto/pb_experiment"
	"context"
)

func main() {
	experimentClient := client.NewExperimentClient()
	{
		experimentClient.CreateExperiment(context.Background(),
			&pb.CreateExperimentRequest{
				ExperimentInfo: &pb.ExperimentInfo{
					Name:        "experiment1",
					Description: "experiment1 description",
					Groups: []*pb.ExperimentGroup{
						{
							Name:                   "group1",
							Description:            "group1 description",
							Allocation:             0.5,
							FromExperimentId:       0,
							WhitelistUserpackageID: 0,
						},
						{
							Name:                   "group2",
							Description:            "group2 description",
							Allocation:             0.5,
							FromExperimentId:       0,
							WhitelistUserpackageID: 0,
						},
					},
					Status: pb.ExperimentStatus_EXPERIMENT_STATUS_READYING,
				},
			},
		)
	}
}
