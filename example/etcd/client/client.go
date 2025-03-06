package main

import (
	"context"
	"log"
	"time"

	pb "etcd_demo/pb/hello" // 根据实际情况修改包路径
	resolve "etcd_demo/resolve"

	clientv3 "go.etcd.io/etcd/client/v3"
	"google.golang.org/grpc"
	"google.golang.org/grpc/resolver"
)

func main() {
	etcdClient, err := clientv3.New(clientv3.Config{
		Endpoints: []string{"localhost:2379"},
	})
	if err != nil {
		log.Fatal(err)
	}
	defer etcdClient.Close()

	resolver.Register(&resolve.EtcdResolverBuilder{EtcdClient: etcdClient})

	// 使用自定义解析器创建 gRPC 连接
	serviceConfig := `{"loadBalancingPolicy": "round_robin"}`
	conn, err := grpc.Dial("etcd:///greeter-service",
		grpc.WithInsecure(),
		grpc.WithDefaultServiceConfig(serviceConfig))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewGreeterClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	r, err := c.SayHello(ctx, &pb.HelloRequest{Name: "World"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetMessage())
}
