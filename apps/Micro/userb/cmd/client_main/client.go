package main

import (
	client "ABTest/apps/Micro/userb/client"
	pb "ABTest/pkgs/proto/pb_userb"
	"context"
	"fmt"

	"google.golang.org/protobuf/proto"
)

func main() {
	userbClient := client.NewUserbClient()
	{
		// 编写测试样例
		resp, err := userbClient.GetUserInfo(context.Background(), &pb.GetUserInfoRequest{
			UserInfo: &pb.UserInfo{
				UserId: proto.Uint64(1),
			},
		})
		if err != nil {
			panic(err)
		}
		if resp == nil {
			fmt.Printf("%v", resp)
		}
	}
}
