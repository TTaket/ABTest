package main

import (
	client "ABTest/apps/Micro/layer/client"
	"ABTest/pkgs/proto/pb_layer"
	"context"
	"fmt"
)

func main() {
	layerClient := client.NewLayerClient()
	{
		// 编写测试样例
		resp, err := layerClient.GetLayer(context.Background(), &pb_layer.GetLayerRequest{
			Id: "1",
		},
		)
		if err != nil {
			panic(err)
		}
		if resp == nil {
			fmt.Printf("%v", resp)
		}
	}
}
