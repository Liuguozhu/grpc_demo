package main

import (
	"context"
	"google.golang.org/grpc"
	"grpc_demo/src/pb"
	"log"
	"net"
)

type server struct {
	//pb.UnimplementedGreeterServer
}

func (s *server) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloReply, error) {
	message := "Hello " + req.GetName() + ", this is grpc go server response!"
	return &pb.HelloReply{Message: message}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	log.Println("Server started, listening on 50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
