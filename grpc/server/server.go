package server

import (
	"context"
	"log"
	"strings"

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

func (s *HelloServer) SayHelloSpecific(ctx context.Context, req *pb.RepeatedHelloRequest) (*pb.HelloReply, error) {
	var result string
	myName := "rushikesh"

	for _, name := range req.GetNames() {
		if res := strings.Compare(myName, name); res == 0 {
			result = myName
			break
		} else {
			result = "NotDefined"
		}
	}

	return &pb.HelloReply{Message: "Hello " + result}, nil
}
