package main

import (
	"fmt"
	"log"
	"net/http"
	"ws-noughts-and-crosses/internal/hub"
)

var mux = http.NewServeMux()

func main() {
	centralHub := hub.NewHub()
	fmt.Println(centralHub)

	go centralHub.Run()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			registration(centralHub, w, r)
	})

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}