package main

import (
	client "ABTest/apps/Xxx/client"
	pb "ABTest/pkgs/proto/pb_xxx"
	"context"
	"fmt"
)

func main() {
	XxxClient := client.NewXxxClient()
	{
		// 请求体
		req := &pb.XxxRequest{}
		// 调用函数
		resp, err := xxxClient.funcname(context.Background(), req)
		if err != nil {
			panic(err)
		}
		fmt.Println(resp.String())
	}
}
