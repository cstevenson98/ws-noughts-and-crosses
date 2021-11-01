package handlers

import (
	"log"
	"net/http"
	"ws-noughts-and-crosses/internal/hub"
)

func Registration(centralHub *hub.Hub, w http.ResponseWriter, r *http.Request) {

	log.Print("Starting client and assigning game")

	Upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	conn, err := Upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	client := &hub.Player{Hub: centralHub, Conn: conn, Stream: make(chan []byte, 1)}
	centralHub.Register <- client

	go client.ReadPump()
	go client.WritePump()
}
