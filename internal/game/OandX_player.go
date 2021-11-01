package game

import (
	"fmt"
	"github.com/gorilla/websocket"
	"ws-noughts-and-crosses/internal/hub"
)

type NoughtsAndCrossesPlayer struct {
	IPlayer
	Hub    *hub.Hub
	Game   NoughtsAndCrossesGame
	Conn   *websocket.Conn
	Stream chan []byte
}

func (x NoughtsAndCrossesPlayer) ReadPump() {
	return
}

func (x NoughtsAndCrossesPlayer) WritePump() {
	return
}

func (x NoughtsAndCrossesPlayer) GetGame() IGame {
	return x.Game
}

func (x NoughtsAndCrossesPlayer) MessagePlayer(message []byte) {
	fmt.Println(string(message))
	return
}