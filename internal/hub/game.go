package hub

import (
	"encoding/json"
	"fmt"
	"time"
)

const (
	GameWaiting = "Waiting"
	MaxPlayers  = 20
	dtDefault   = time.Millisecond * 50
	PlayerSpeed = 0.003 // px/ms
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
			var actions []PlayerActionMessage
			for otherPlayer := range g.Players {
				if otherPlayer != player {
					actions = append(actions, PlayerActionMessage{
						ID:        otherPlayer.ID,
						Direction: otherPlayer.Direction,
						Pos:       otherPlayer.Pos,
						Vel:       otherPlayer.Vel,
					})
				}
			}

			message := CurrentStateMessage{
				Update: actions,
			}
			payload, _ := json.Marshal(message)
			player.Stream <- payload
		}
		g.t0 += g.dt.Seconds()
	}
}
