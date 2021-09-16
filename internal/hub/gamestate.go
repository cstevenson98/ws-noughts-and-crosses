package hub

import "fmt"

const(
	GameStandby = 1
	GamePlayer1 = 2
	GamePlayer2 = 3
	GameOver    = 4
)

type GameState struct {
	Board     [3][3]string
	TurnState int
}

func (g *GameState) SetSquare(x, y int, kind string) error {
	// Should only be called if the player is validated,
	// *and* the position is available
	if g.Board[x][y] != "" {
		return fmt.Errorf("cannot write to %d, %d, already occupied", x, y)
	}
	g.Board[x][y] = kind
	return nil
}

func (g *GameState) Clear() {
	g.Board = [3][3]string{{"", "", ""}, {"", "", ""}, {"", "", ""}}
}