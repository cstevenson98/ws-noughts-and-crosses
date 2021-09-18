package hub

import (
	"github.com/gorilla/websocket"
	"log"
)

type Client struct {
	Hub          *Hub
	Conn         *websocket.Conn
}

func (c *Client) ReadPump() {
	for {
		_, _, err := c.Conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}
	}
}

func (c *Client) WritePump() {
	return
}
