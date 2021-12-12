package main

import (
	"log"
	"net/http"
	"time"
	"ws-noughts-and-crosses/internal/hub"
)

func main() {
	centralHub := hub.NewHub()
	go centralHub.Run()
	go centralHub.LogOnInterval(time.Second * 30)

	initRouter(centralHub)

	err := http.ListenAndServe(":8765", mux)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}