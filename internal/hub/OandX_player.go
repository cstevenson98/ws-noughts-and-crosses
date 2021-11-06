package hub

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
)

type OandXPlayer struct {
	IPlayer

	Hub  *Hub
	Game *OandXGame
	Conn *websocket.Conn
	Stream chan []byte
}

func (p *OandXPlayer) ReadPump() {
	defer func() {
		p.Hub.Unregister <- p
		p.Conn.Close()
	}()

	for {
		_, message, err := p.Conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}

		p.Hub.MakeTurn <- Turn{p, message}
	}
}

func (p *OandXPlayer) WritePump() {
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


func (p *OandXPlayer) ProcessTurn(turn []byte) error {
	var turnCoords [2]int
	jsonErr := json.Unmarshal(turn, &turnCoords)
	if jsonErr != nil {
		log.Printf("ERROR: unmarshalling")
		return fmt.Errorf("unable to unmarshal turn []byte(%q)", string(turn))
	}


	var nextBoard = p.Hub.OandXGames[p.Game]
	playerLabel := p.Game.WhichPlayer(p)
	var otherPlayer *OandXPlayer
	if playerLabel == GamePlayer1 && p.Game.Status == GamePlayer1 {
		nextBoard.Board[turnCoords[0]][turnCoords[1]] = "X"
		otherPlayer = p.Game.Player2
		p.Game.Status = GamePlayer2
	} else if playerLabel == GamePlayer2 && p.Game.Status == GamePlayer2 {
		nextBoard.Board[turnCoords[0]][turnCoords[1]] = "0"
		otherPlayer = p.Game.Player1
		p.Game.Status = GamePlayer1
	} else {
		return nil
	}

	output := nextBoard.BoardToOutput()
	p.Stream <- output
	if otherPlayer != nil {
		otherPlayer.Stream <- output
	}
	p.Hub.OandXGames[p.Game] = nextBoard
	return nil
}