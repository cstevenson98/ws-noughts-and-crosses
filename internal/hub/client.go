package hub

import (
	"github.com/gorilla/websocket"
)

type Client struct {
	GameNumber   int
	PlayerNumber int
	State        GameState
	Hub          *Hub
	Conn         *websocket.Conn
}

