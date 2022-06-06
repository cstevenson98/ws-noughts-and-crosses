package hub

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"time"
)

const (
	GameWaiting        = "Waiting"
	MaxPlayers         = 20
	dtDefault          = time.Millisecond * 1000
	subDt              = time.Millisecond * 100
	DragConstant       = 0.1
	PlayerAcceleration = 0.001
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
			message := CurrentStateMessage{
				MyPosition: [2]float64{0.5 + 0.25*math.Sin(g.t0), 0.5 + 0.25*math.Cos(g.t0)},
				Positions:  [][2]float64{},
			}
			payload, _ := json.Marshal(message)
			player.Stream <- payload

			log.Println(player)
			player.InputStack.Reset()
		}
		g.t0 += g.dt.Seconds()
	}
}
