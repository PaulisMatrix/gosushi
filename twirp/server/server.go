package main

import (
	"context"
	"fmt"
	"net/http"

	pb "practice/twirp/ratelimit"
	"practice/twirp/store"
)

type RatelimitServer struct{}

func (s *RatelimitServer) GetRatelimitInfo(ctx context.Context, req *pb.GetRatelimitRequest) (*pb.GetRatelimitResponse, error) {
	return nil, nil
}

func (s *RatelimitServer) SetRatelimitInfo(ctx context.Context, req *pb.SetRatelimitRequest) (*pb.SetRatelimitResponse, error) {
	Ratelimit := store.RateLimit{
		ClientID:  req.GetClientID(),
		RouterMap: store.UnMarshalProtoToMap(req.GetRouterMap()),
	}

	fmt.Println("Unmarshalled RouterMap:")
	for _, m := range Ratelimit.RouterMap {
		fmt.Printf("Rate %d and Endpoint %s\n", m.Rate, m.Endpoint)
	}

	return &pb.SetRatelimitResponse{}, nil
}

func main() {
	twirpHandler := pb.NewRatelimitServer(&RatelimitServer{})
	// You can use any mux you like - NewHelloWorldServer gives you an http.Handler.
	mux := http.NewServeMux()
	// The generated code includes a method, PathPrefix(), which
	// can be used to mount your service on a mux.
	fmt.Println("Server listening on port 8080")
	mux.Handle(twirpHandler.PathPrefix(), twirpHandler)
	http.ListenAndServe(":8080", mux)

}
