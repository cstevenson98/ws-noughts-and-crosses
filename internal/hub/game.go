package hub

import (
	"encoding/json"
	"fmt"
	"time"
	"ws-noughts-and-crosses/pkg/vec"
)

const (
	GameWaiting = "Waiting"
	MaxPlayers  = 20
	dtDefault   = time.Millisecond * 100
	PlayerSpeed = 0.0003 // px/ms
)

type Turn struct {
	player      IPlayer
	encodedTurn []byte
}

type IState interface{}

type IPlayer interface {
	ReadPump()
	WritePump()
}

type IGame interface{}

type Game struct {
	IGame
	Players map[*Player]bool
	Status  string
	t0      float64
	dt      time.Duration
}

func (g *Game) AddClient(player *Player) error {
	if _, exists := g.Players[player]; exists {
		return fmt.Errorf("player already exists")
	}

	// Add player
	g.Players[player] = true

	player.Game = g
	return nil
}

func (g *Game) SlotsFree() int {
	return MaxPlayers - len(g.Players)
}

func (g *Game) RunGame() {
	for range time.Tick(g.dt) {
		for player := range g.Players {
			now := float64(time.Now().UnixMilli())
			player.Evolve(now)

			otherPositions := make([]vec.Vec, 0)
			for otherPlayer := range g.Players {
				if otherPlayer != player {
					otherPositions = append(otherPositions, otherPlayer.Pos)
				}
			}

			message := CurrentStateMessage{
				MyPosition: player.Pos,
				Positions:  otherPositions,
			}
			payload, _ := json.Marshal(message)
			player.Stream <- payload
		}
		g.t0 += g.dt.Seconds()
	}
}
