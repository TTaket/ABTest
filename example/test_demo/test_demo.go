package testdemo

import (
	"context"
	"fmt"

	"github.com/TTaket/ABTest"

	"google.golang.org/protobuf/proto"
)

// TestDemo 测试示例
func TestDemo() {
	// 初始化UserbClient
	userbClient := ABTest.NewUserbClient()
	// 初始化ExperimentClient
	experimentClient := ABTest.NewExperimentClient()

	// 调用UserbClient的AddUser方法
	{
		resp, err := userbClient.GetUserInfo(context.Background(), &ABTest.pb_userb.GetUserInfoRequest{
			UserInfo: &ABTest.pb_userb.UserInfo{
				UserId: proto.Uint64(1),
			},
		})
		if err != nil {
			panic(err)
		}
		fmt.Printf("TestDemo: get user info %s", resp.String())
	}
	// 调用ExperimentClient的AddExperiment方法
	{
		resp, err := experimentClient.GetExperiment(context.Background(), &ABTest.pb_experiment.GetExperimentRequest{
			ExperimentId: "1",
		})
		if err != nil {
			panic(err)
		}
		fmt.Printf("TestDemo: get experiment %s", resp.String())
	}
}
