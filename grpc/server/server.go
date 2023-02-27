package server

import (
	"context"
	"log"

	pb "practice/grpc/hello_world"
)

// server is used to implement helloworld.GreeterServer.
type HelloServer struct {
	pb.UnimplementedGreeterServer
}

// SayHello implements helloworld.GreeterServer
func (s *HelloServer) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Got request for: %v", req.GetName())
	return &pb.HelloReply{Message: "Hello " + req.GetName()}, nil
}
