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

// NewState creates a blank GameState.
func NewState() *GameState {
	return &GameState{[3][3]string{{" ", " ", " "}, {" ", " ", " "}, {" ", " ", " "}}, GameStandby}
}

// BoardToOutput outputs the noughts-and-crosses state as a user-
// readable ASCII art
func (g *GameState) BoardToOutput() []byte {
	var outString string
	outString += g.Board[0][0]+"|"+g.Board[1][0]+"|"+g.Board[2][0]+"\n"
	outString += "-----\n"
	outString += g.Board[0][1]+"|"+g.Board[1][1]+"|"+g.Board[2][1]+"\n"
	outString += "-----\n"
	outString += g.Board[0][2]+"|"+g.Board[1][2]+"|"+g.Board[2][2]+"\n"
	return []byte(outString)
}