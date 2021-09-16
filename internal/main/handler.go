package main

import (
	"log"
	"net/http"
	"ws-noughts-and-crosses/internal/hub"
)

func registration(hub *hub.Hub, w http.ResponseWriter, r *http.Request) {
	log.Print("Starting client and assigning game")
}

