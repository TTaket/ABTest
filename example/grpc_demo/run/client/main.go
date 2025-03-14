package main

import (
	"context"
	"fmt"
	"grpcdemo/proto/server1/pb_server1/pb_server1"
	"grpcdemo/proto/server2/pb_server2/pb_server2"
	"grpcdemo/servers/server1"
	"grpcdemo/servers/server2"
)

func main() {
	{
		r, err := server1.NewServer1Client().Get(context.Background(), &pb_server1.GetRequest{
			ExperimentId: "1",
		})
		if err != nil {
			panic(err)
		}
		fmt.Println(r.String())
	}
	{
		r, err := server2.NewServer2Client().Get(context.Background(), &pb_server2.GetRequest{
			ExperimentId: "1",
		})
		if err != nil {
			panic(err)
		}
		fmt.Println(r.String())
	}
}
