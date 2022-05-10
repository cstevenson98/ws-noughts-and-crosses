package hub

import (
	"fmt"
)

type GameState struct {
	IState
	Board [3][3]string
}

// SetSquare takes a pair of integers and the token and modifies the
// board at that location.
func (g *GameState) SetSquare(x, y int, kind string) error {
	// Should only be called if the player is validated,
	// *and* the position is available
	if g.Board[x][y] != "" {
		return fmt.Errorf("cannot write to %d, %d, already occupied", x, y)
	}
	g.Board[x][y] = kind
	return nil
}

// NewState creates a blank GameState.
func NewState() *GameState {
	return &GameState{Board: [3][3]string{{" ", " ", " "}, {" ", " ", " "}, {" ", " ", " "}}}
}

func (g *GameState) GameStateOutput() (out []byte) {
	return
}
