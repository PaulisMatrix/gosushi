package main

import (
	"context"
	"fmt"
	"net/http"

	store "practice/twirp"
	pb "practice/twirp/ratelimit"
)

func main() {
	client := pb.NewRatelimitProtobufClient("http://localhost:8080", &http.Client{})

	//_, err := client.GetRatelimitInfo(context.Background(), &pb.GetRatelimitRequest{ClientID: "sebmulti21"})

	setRequest := &pb.SetRatelimitRequest{
		ClientID:  "sebmulti21",
		RouterMap: store.ConvertToProto(store.GetRouterMap()),
	}
	_, err := client.SetRatelimitInfo(context.Background(), setRequest)

	if err == nil {
		fmt.Println("RouterMAP Updated: ")
		for _, m := range store.RateLimitInfo.RouterMap {
			fmt.Printf("Endpoint is %s and rate is %d\n", m.Endpoint, m.Rate)
		}
	} else {
		fmt.Println("ERROR: ", err)
	}
}
