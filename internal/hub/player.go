package hub

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"log"
	"ws-noughts-and-crosses/pkg/vec"
)

// Player Stores information about and facilitates communication
// with client players over websocket.
type Player struct {
	IPlayer
	InputStack InputStack

	Pos vec.Vec
	Vel vec.Vec

	Hub    *Hub
	Game   *Game
	Conn   *websocket.Conn
	Stream chan []byte
}

func (p *Player) ReadPump() {
	defer func() {
		p.Hub.Unregister <- p
		p.Conn.Close()
	}()

	for {
		_, msg, err := p.Conn.ReadMessage()
		if err != nil {
			if websocket.IsCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}

		var input UserInputEventMessage
		err = json.Unmarshal(msg, &input)
		if err != nil {
			log.Println("could not unmarshal user input")
		}

		p.InputStack.Push(input)
	}
}

func (p *Player) WritePump() {
	defer func() {
		p.Conn.Close()
	}()

	for {
		select {
		case message, ok := <-p.Stream:
			if !ok {
				// The hub closed the channel.
				p.Conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			w, err := p.Conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}
			w.Write(message)

			if err := w.Close(); err != nil {
				return
			}
		}
	}
}
