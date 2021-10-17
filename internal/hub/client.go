package hub

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"log"
)

var (
	newline = []byte{'\n'}
	space   = []byte{' '}
)

type Player struct {
	Hub          *Hub
	Game         *Game
	Conn         *websocket.Conn
	Stream       chan []byte
}

func (c *Player) ReadPump() {
	defer func() {
		c.Hub.Unregister <- c
		c.Conn.Close()
	}()

	for {
		_, message, err := c.Conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}

		var turnCoords [2]int
		jsonErr := json.Unmarshal(message, &turnCoords)
		if jsonErr != nil {
			log.Printf("ERROR: unmarshalling")
			break
		}

		c.Hub.MakeTurn <- Turn{ c, turnCoords }
	}
}

func (c *Player) WritePump() {
	defer func() {
		c.Conn.Close()
	}()

	for {
		select {
		case message, ok := <-c.Stream:
			if !ok {
				// The hub closed the channel.
				c.Conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			w, err := c.Conn.NextWriter(websocket.TextMessage)
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
