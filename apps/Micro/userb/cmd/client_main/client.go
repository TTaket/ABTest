package main

import (
	client "ABTest/apps/Micro/userb/client"
	pb "ABTest/pkgs/proto/pb_userb"
	"context"
	"fmt"

	"google.golang.org/protobuf/proto"
)

func main() {
	UserbClient := client.NewUserbClient()
	{
		// 请求体
		req := &pb.AddUserRequest{
			UserInfo: &pb.UserInfo{
				Name: proto.String("test"),
			},
		}
		// 调用函数
		resp, err := UserbClient.AddUser(context.Background(), req)
		if err != nil {
			panic(err)
		}
		fmt.Println(resp.String())
	}
}
