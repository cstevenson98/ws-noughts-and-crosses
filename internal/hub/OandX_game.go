package hub

import (
	"encoding/json"
	"fmt"
	"math"
	"time"
)

const MaxPlayers = 20
const dtDefault = time.Millisecond * 100

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
			message := CurrentStateMessage{
				Positions: [][2]float64{{0.5 + 0.25*math.Sin(g.t0), 0.5 + 0.25*math.Cos(g.t0)}},
			}
			payload, _ := json.Marshal(message)
			player.Stream <- payload
		}
		g.t0 += g.dt.Seconds()
	}
}