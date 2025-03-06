package main

import (
	"context"
	"log"
	"net"

	pb "etcd_demo/pb/hello" // 根据实际情况修改包路径

	clientv3 "go.etcd.io/etcd/client/v3"
	"google.golang.org/grpc"
)

// 实现 Greeter 服务
type server struct {
	pb.UnimplementedGreeterServer
}

func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &pb.HelloReply{Message: "Hello " + in.GetName()}, nil
}

func registerService(etcdClient *clientv3.Client, serviceName, instanceAddr string, ttl int64) error {
	leaseResp, err := etcdClient.Grant(context.TODO(), ttl)
	if err != nil {
		return err
	}
	leaseID := leaseResp.ID

	keepAliveChan, err := etcdClient.KeepAlive(context.TODO(), leaseID)
	if err != nil {
		return err
	}
	go func() {
		for range keepAliveChan {
			// 续租成功，可添加日志等操作
		}
	}()

	key := "services/" + serviceName + "/" + instanceAddr
	_, err = etcdClient.Put(context.TODO(), key, instanceAddr, clientv3.WithLease(leaseID))
	return err
}

func main() {
	etcdClient, err := clientv3.New(clientv3.Config{
		Endpoints: []string{"localhost:2379"},
	})
	if err != nil {
		log.Fatal(err)
	}
	defer etcdClient.Close()

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})

	serviceName := "greeter-service"
	instanceAddr := "localhost:50051"
	ttl := int64(10)

	err = registerService(etcdClient, serviceName, instanceAddr, ttl)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
