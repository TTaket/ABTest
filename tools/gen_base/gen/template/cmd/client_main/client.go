package main

import (
	client "ABTest/apps/Micro/%1/client"
	pb "ABTest/pkgs/proto/pb_%1"
	"context"
	"fmt"
)

func main() {
	%3Client := client.New%2Client()
	{
		// 编写测试样例		
	}
}
