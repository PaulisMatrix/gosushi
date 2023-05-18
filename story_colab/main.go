package storycolab

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"practice/story_colab/config"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Healthcheck route!!")
	if r.URL.Path == "/" {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("Starting Story Colab Server!!"))
}

func StoryColab() {
	mux := http.NewServeMux()
	mux.HandleFunc("/ping", HealthCheck)
	conn_str := config.GetConfig().MongoURI
	Connect(conn_str)

	fmt.Println("Starting the server!!!")
	s := &http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	log.Fatal(s.ListenAndServe())
}

func Connect(mongoURI string) {
	// connecting to the cluster
	client, err := mongo.NewClient(options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	// list databases
	databases, err := client.ListDatabaseNames(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Available databases", databases)
}
