package twirp

import pb "practice/twirp/ratelimit"

const DefaultRate = 10

type RouterMap struct {
	Endpoint string `json:"endpoint"`
	Rate     int    `json:"rate"`
}

//default routermapping
//continue with the map as we add more endpoints.
//var RM = &[]RouterMap{{Endpoint: "ping", Rate: DefaultRate}, {Endpoint: "business_hour_groups", Rate: DefaultRate},
//	{Endpoint: "callback_central", Rate: DefaultRate},
//}

type RateLimit struct {
	ClientID  string       `json:"client_id" bson:"client_id"`
	RouterMap []*RouterMap `json:"routers_map_list" bson:"routers_map_list"`
}

var RateLimitInfo = RateLimit{
	ClientID:  "sebmulti21",
	RouterMap: GetRouterMap(),
}

func GetRouterMap() []*RouterMap {
	var rmap []*RouterMap

	rmap = append(rmap, &RouterMap{
		Endpoint: "ping",
		Rate:     DefaultRate,
	})

	rmap = append(rmap, &RouterMap{
		Endpoint: "business_hour_groups",
		Rate:     DefaultRate,
	})

	return rmap

}

func ConvertToProto(rmap []*RouterMap) []*pb.RouterMap {
	var result []*pb.RouterMap

	for _, m := range rmap {
		temp := &pb.RouterMap{
			Endpoint: m.Endpoint,
			Rate:     int32(m.Rate),
		}
		result = append(result, temp)
	}

	return result
}
