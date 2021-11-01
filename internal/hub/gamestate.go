package hub

import "fmt"

type GameState struct {
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
	return &GameState{[3][3]string{{" ", " ", " "}, {" ", " ", " "}, {" ", " ", " "}}}
}

func (g *GameState) BoardToOutput() []byte {
	var outString string
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			toConcatenate := g.Board[i][j]
			outString = outString + toConcatenate
		}
	}
	return []byte(outString)
}

// BoardToString outputs the noughts-and-crosses state as a user-
// readable ASCII art
func (g *GameState) BoardToString() string {
	var outString string
	outString += g.Board[0][0] + "|" + g.Board[1][0] + "|" + g.Board[2][0] + "\n"
	outString += "-----\n"
	outString += g.Board[0][1] + "|" + g.Board[1][1] + "|" + g.Board[2][1] + "\n"
	outString += "-----\n"
	outString += g.Board[0][2] + "|" + g.Board[1][2] + "|" + g.Board[2][2] + "\n"
	return outString
}

// Clear resets the board.
func (g *GameState) Clear() {
	g.Board = [3][3]string{{" ", " ", " "}, {" ", " ", " "}, {" ", " ", " "}}
	return
}
