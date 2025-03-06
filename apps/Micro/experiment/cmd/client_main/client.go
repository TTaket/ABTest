package main

import (
	client "ABTest/apps/Micro/experiment/client"
	pb "ABTest/pkgs/proto/pb_experiment"
	"context"
	"fmt"
)

func main() {
	experimentClient := client.NewExperimentServiceClient()
	// get experiment
	{
		req := &pb.GetExperimentRequest{
			ExperimentId: "1",
		}
		resp, err := experimentClient.GetExperiment(context.Background(), req)
		if err != nil {
			panic(err)
		}
		fmt.Println(resp.String())
	}
}
