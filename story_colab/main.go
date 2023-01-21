package storycolab

import (
	"log"
	"net/http"
	"time"
)

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Starting Story Colab Server!!"))
}

func StoryColab() {
	mux := http.NewServeMux()
	mux.HandleFunc("/ping", HealthCheck)

	s := &http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	log.Fatal(s.ListenAndServe())
}
