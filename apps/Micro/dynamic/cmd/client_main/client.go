package main

import (
	client "ABTest/apps/Micro/dynamic/client"
	pb "ABTest/pkgs/proto/pb_dynamic"
	"context"
	"fmt"
)

func main() {
	dynamicClient := client.NewDynamicClient()
	{
		// 编写测试样例
		resp, err := dynamicClient.GetExperimentLayerBinding(context.Background(), &pb.GetExperimentLayerBindingRequest{
			ExperimentId: "123",
			LayerId:      "456",
		})
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		fmt.Println("Response:", resp)
	}
}
