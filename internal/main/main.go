package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"time"
	"ws-noughts-and-crosses/internal/hub"
)

var mux = http.NewServeMux()
var upgrader = websocket.Upgrader{}

func main() {
	centralHub := hub.NewHub()
	fmt.Println(centralHub)

	go centralHub.Run()
	go centralHub.LogOnInterval(time.Second * 30)

	fileServer := http.FileServer(http.Dir("./frontend"))
	mux.Handle("/", fileServer)

	mux.HandleFunc("/connect",
		func(w http.ResponseWriter, r *http.Request) {
			registration(centralHub, w, r)
	})


	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
