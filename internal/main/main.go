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
	go centralHub.LogOnInterval(time.Second * 1)

	initRouter(centralHub)

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
