package main

import (
	client "ABTest/apps/Micro/experiment/client"
	pb "ABTest/pkgs/proto/pb_experiment"
	"context"
	"fmt"
)

func main() {
	experimentClient := client.NewExperimentClient()
	// get experiment demo
	{
		req := &pb.GetExperimentRequest{
			ExperimentId: "1",
		}
		resp, err := experimentClient.GetExperiment(context.Background(), req)
		if err != nil {
			// 用户视角 不能使用mylogger
			fmt.Println(err)
		}
		fmt.Println("用户行为...", resp)
	}
}
