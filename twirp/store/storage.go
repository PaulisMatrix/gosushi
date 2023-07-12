package store

import (
	pb "practice/twirp/ratelimit"
)

type Endpoint string

const (
	Ping   Endpoint = "ping"
	Pong   Endpoint = "pong"
	Unkown Endpoint = "unkown"
)

type RouterMap struct {
	Endpoint string `json:"endpoint"`
	Rate     int    `json:"rate"`
}

//default routermapping
//continue with the map as we add more endpoints.
var RM = []*RouterMap{{Endpoint: string(Ping), Rate: 10}, {Endpoint: string(Pong), Rate: 20}}

type RateLimit struct {
	ClientID  string       `json:"client_id" bson:"client_id"`
	RouterMap []*RouterMap `json:"routers_map_list" bson:"routers_map_list"`
}

func MarshalEndpointToProtoEnum(endpoint Endpoint) pb.RouterMap_APIEndpoint {
	switch endpoint {
	case Ping:
		return pb.RouterMap_PING
	case Pong:
		return pb.RouterMap_PONG
	default:
		return pb.RouterMap_UNKNOWN
	}

}

func UnMarshalEndpointToEnum(t pb.RouterMap_APIEndpoint) string {
	switch t {
	case pb.RouterMap_PING:
		return string(Ping)
	case pb.RouterMap_PONG:
		return string(Pong)
	default:
		return string(Unkown)
	}

}
func MarshalMapToProto(rmap []*RouterMap) []*pb.RouterMap {
	var result []*pb.RouterMap

	for _, m := range rmap {
		APIEndpoint := Endpoint(m.Endpoint)

		temp := &pb.RouterMap{
			Rate:     int32(m.Rate),
			Endpoint: MarshalEndpointToProtoEnum(APIEndpoint),
		}
		result = append(result, temp)
	}

	return result
}

func UnMarshalProtoToMap(rmap []*pb.RouterMap) []*RouterMap {
	var result []*RouterMap

	for _, m := range rmap {
		temp := &RouterMap{
			Rate:     int(m.GetRate()),
			Endpoint: UnMarshalEndpointToEnum(m.GetEndpoint()),
		}

		result = append(result, temp)
	}

	return result
}
