package testdemo

import (
	client2 "ABTest/apps/Micro/experiment/client"
	client1 "ABTest/apps/Micro/userb/client"
	"ABTest/pkgs/proto/pb_experiment"
	"ABTest/pkgs/proto/pb_userb"
	"context"
	"fmt"

	"google.golang.org/protobuf/proto"
)

// DoubleModule 测试示例
func DoubleModule() {
	// 初始化UserbClient
	userbClient := client1.NewUserbClient()
	// 初始化ExperimentClient
	experimentClient := client2.NewExperimentClient()

	// 调用UserbClient的AddUser方法
	{
		resp, err := userbClient.GetUserInfo(context.Background(), &pb_userb.GetUserInfoRequest{
			UserInfo: &pb_userb.UserInfo{
				UserId: proto.Uint64(1),
			},
		})
		if err != nil {
			panic(err)
		}
		fmt.Printf("TestDemo: get user info %s\n", resp.String())
	}
	// 调用ExperimentClient的AddExperiment方法
	{
		resp, err := experimentClient.GetExperiment(context.Background(), &pb_experiment.GetExperimentRequest{
			ExperimentId: "1",
		})
		if err != nil {
			panic(err)
		}
		fmt.Printf("TestDemo: get experiment %s\n", resp.String())
	}
}
