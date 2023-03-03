package main

import (
	"context"
	"log"
	"time"

	pb "practice/grpc/hello_world"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const serverAddress string = ":4772"

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(serverAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	//r, err := c.SayHello(ctx, &pb.HelloRequest{Name: "rushikesh"})
	names := []string{"albin", "nithya", "rushikesh"}
	r, err := c.SayHelloSpecific(ctx, &pb.RepeatedHelloRequest{Names: names})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetMessage())
}
