package hub

import (
	"fmt"
)

type OandXGame struct {
	IGame
	Player1 *OandXPlayer
	Player2 *OandXPlayer
	Status  string
}

// AddClient adds a client to a game in the first available slot and tells player
// what game they are in by assigning it to Game field.
func (g *OandXGame) AddClient(player *OandXPlayer) error {
	if g.Player1 == nil {
		g.Player1 = player
	} else if g.Player2 == nil {
		g.Player2 = player
		g.Status = GamePlayer1
	} else {
		return fmt.Errorf("no slots in game")
	}
	player.Game = g
	return nil
}

// SlotsFree returns the number of unoccupied slots in a game.
func (g *OandXGame) SlotsFree() (slots int) {
	if g.Player1 == nil {
		slots += 1
	}
	if g.Player2 == nil {
		slots += 1
	}
	return
}

// WhichPlayer determines if player is in slot 1 or slot 2.
func (g *OandXGame) WhichPlayer(player *OandXPlayer) string {
	if player == g.Player1 {
		return GamePlayer1
	} else if player == g.Player2 {
		return GamePlayer2
	}
	return "<Unknown Player>"
}