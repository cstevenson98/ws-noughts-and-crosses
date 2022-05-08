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

	mux.HandleFunc("/noughtsAndCrosses",
		func(w http.ResponseWriter, r *http.Request) {
			handlers.NoughtsAndCrosses(w, r)
		})

	mux.HandleFunc("/noughtsAndCrosses/connect",
		func(w http.ResponseWriter, r *http.Request) {
			handlers.OandXRegistration(centralHub, w, r)
		})

	mux.HandleFunc("/login",
		func(w http.ResponseWriter, r *http.Request) {
			centralHub.Login(w, r)
		})

	mux.HandleFunc("/signup",
		func(w http.ResponseWriter, r *http.Request) {
			centralHub.Signup(w, r)
		})

	fileServer := http.FileServer(http.Dir("./frontend/"))
	mux.Handle("/frontend/", http.StripPrefix("/frontend", fileServer))
}
