package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func home(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "websockets.html")
}

func echo(w http.ResponseWriter, r *http.Request) {

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("Http connection upgrade to websockets failed:", err)
	}
	defer conn.Close()

	for {

		//Read the message from the browser
		msgType, msg, err := conn.ReadMessage()
		if err != nil {
			log.Print("Unable to read the message")
			return
		}

		//Print the message to console
		fmt.Printf("%s sent: %s\n", conn.RemoteAddr(), string(msg))

		//Write the message back to the browser
		if err := conn.WriteMessage(msgType, msg); err != nil {
			log.Print("Unable to write back message")

		}

	}

}
func main() {
	http.HandleFunc("/echo", echo)
	http.HandleFunc("/", home)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
