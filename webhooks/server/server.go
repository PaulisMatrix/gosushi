package server

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func myHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello There!!"))

}

func router() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/ping", myHandler)
	return mux
}

func Serve() {
	mux := router()
	fmt.Println("Starting my webhooks server!!")
	s := &http.Server{
		Addr:         ":8080",
		Handler:      mux,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(s.ListenAndServe())
}
