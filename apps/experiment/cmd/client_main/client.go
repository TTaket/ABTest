package main

import (
	"context"
	"flag"
	"log"
	"time"

	pb "ABTest/pkgs/proto/pb_experiment"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	defaultName = "world"
)

var (
	addr = flag.String("addr", "localhost:50051", "the address to connect to")
	name = flag.String("name", defaultName, "Name to greet")
)

func main() {
	flag.Parse()
	// Set up a connection to the server.
	conn, err := grpc.NewClient(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewExperimentServiceClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.GetExperiment(ctx, &pb.GetExperimentRequest{ExperimentId: "1"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("mock: %s", r.String())
}
