package main

import (
	"log"
	"net/http"
	"ws-noughts-and-crosses/internal/hub"
)

func registration(centralHub *hub.Hub, w http.ResponseWriter, r *http.Request) {
	// First upgrade this connection to a websocket connection,
	// then create a client, add this to the hub.
	log.Print("Starting client and assigning game")
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	client := &hub.Client{Hub: centralHub, Conn: conn}

	// Tell the hub to register this client
	centralHub.Register <- client

	go client.ReadPump()
	go client.WritePump()

	// Launch the read and write pumps for the new client
}

