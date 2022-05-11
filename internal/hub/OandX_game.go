package hub

import "fmt"

const MaxPlayers = 20

type Game struct {
	IGame
	Players map[*Player]bool
	Status  string
	t0      float64
}

func (g *Game) AddClient(player *Player) error {
	if _, exists := g.Players[player]; exists {
		return fmt.Errorf("player already exists")
	}
	g.Players[player] = true // add player

	player.Game = g
	return nil
}

func (g *Game) SlotsFree() int {
	return MaxPlayers - len(g.Players)
}