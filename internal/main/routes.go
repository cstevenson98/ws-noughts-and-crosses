package main

import (
	"net/http"
	"ws-noughts-and-crosses/internal/handlers"
	"ws-noughts-and-crosses/internal/hub"
)

var mux = http.NewServeMux()

func initRouter(centralHub *hub.Hub) {
	mux.HandleFunc("/",
		func(w http.ResponseWriter, r *http.Request) {
			handlers.Home(w, r)
		})

	mux.HandleFunc("/connect",
		func(w http.ResponseWriter, r *http.Request) {
			handlers.Registration(centralHub, w, r)
		})

	fileServer := http.FileServer(http.Dir("./frontend/"))
	mux.Handle("/frontend/", http.StripPrefix("/frontend", fileServer))
}