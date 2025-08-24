package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "github.com/varnit-ta/gRPC-Practice/helloworld"
	"google.golang.org/grpc"
)

type Server struct {
	pb.UnimplementedHelloWorldServiceServer
}

func (s *Server) SayHello(ctx context.Context, in *pb.HelloWorldRequest) (*pb.HelloWorldResponse, error) {
	return &pb.HelloWorldResponse{Message: fmt.Sprintf("My name is %s", in.Name)}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":4000")

	if err != nil {
		log.Fatalf("failed to listend on port 4000: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterHelloWorldServiceServer(s, &Server{})

	log.Printf("gRPC listening at %v", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
