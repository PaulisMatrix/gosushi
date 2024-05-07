package main

import "practice/textsearch"

//"practice/tinycompiler"

//pb "practice/grpc/hello_world"
//"google.golang.org/grpc"
//aoc "practice/adventofcode/day10"
//websockets "practice/websockets"
//cards "practice/cards"
//sc "practice/story_colab"
//server "practice/grpc/server"
//import watcher "practice/snippets"

//const address string = ":4772"

func main() {
	/*
		myset1 := snippets.NewSet()
		myset1.AddMulti([]string{"a", "b", "c", "d"})

		myset2 := snippets.NewSet()
		myset2.AddMulti([]string{"e", "f", "g", "d"})

		inter := myset1.Intersection(myset2)

		inter.Display()

		union := myset1.Union(myset2)

		union.Display()

		diff := myset1.Difference(myset2)

		diff.Display()

		//aoc.SecondPart()
		log.Println("Starting the server!")

		lis, err := net.Listen("tcp", address)
		if err != nil {
			log.Fatal("Failed to Listen on", err)
		}
		log.Println("Listening on", address)

		grpcServer := grpc.NewServer()

		//register your server
		pb.RegisterGreeterServer(grpcServer, &server.HelloServer{})

		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	*/

	// tinycompiler.StartHere()
	textsearch.TextSearch()
	// snippets.Prims()

}
