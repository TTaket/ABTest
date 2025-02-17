package main

import (
	"context"
	"fmt"

	client "ABTest/apps/experiment/client"
	pb "ABTest/pkgs/proto/pb_experiment"
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
