package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	store "practice/twirp"
	pb "practice/twirp/ratelimit"
)

type RatelimitServer struct{}

func (s *RatelimitServer) GetRatelimitInfo(ctx context.Context, req *pb.GetRatelimitRequest) (*pb.GetRatelimitResponse, error) {
	clientID := req.GetClientID()

	if store.RateLimitInfo.ClientID == clientID {
		rMap := store.RateLimitInfo.RouterMap

		resp := &pb.GetRatelimitResponse{
			RouterMap: store.ConvertToProto(rMap),
		}
		return resp, nil
	} else {
		return nil, errors.New("Client ID not found!!")
	}
}

func (s *RatelimitServer) SetRatelimitInfo(ctx context.Context, req *pb.SetRatelimitRequest) (*pb.SetRatelimitResponse, error) {

	rmap := req.GetRouterMap()

	var result []*store.RouterMap

	for _, m := range rmap {
		temp := store.RouterMap{
			Endpoint: m.Endpoint,
			Rate:     int(m.Rate),
		}
		result = append(result, &temp)
	}

	store.RateLimitInfo.RouterMap = append(store.RateLimitInfo.RouterMap, result...)

	return &pb.SetRatelimitResponse{}, nil
}

func main() {
	twirpHandler := pb.NewRatelimitServer(&RatelimitServer{})
	// You can use any mux you like - NewHelloWorldServer gives you an http.Handler.
	mux := http.NewServeMux()
	// The generated code includes a method, PathPrefix(), which
	// can be used to mount your service on a mux.
	mux.Handle(twirpHandler.PathPrefix(), twirpHandler)
	http.ListenAndServe(":8080", mux)
	fmt.Println("Server listening on port 8080")
}
