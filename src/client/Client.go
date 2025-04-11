package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"grpc_demo/src/pb"
	"log"
	"time"
)

func main() {
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewGreeterClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := &pb.HelloRequest{Name: "client_go"}
	reply, err := client.SayHello(ctx, req)
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", reply.GetMessage())
}
