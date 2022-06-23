package handlers

import (
	"github.com/google/uuid"
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

	// Create player and register with central hub
	player := &hub.Player{
		ID:     uuid.New().String(),
		Hub:    centralHub,
		Conn:   conn,
		Stream: make(chan []byte, 1),
		Pos:    [2]float64{0.5, 0.5},
	}
	centralHub.Register <- player

	// Launch goroutine to read/write from client
	go player.ReadPump()
	go player.WritePump()
}
