package hub

import "testing"

// TestPlayer_Evolve1 tests coming back to the same position
func TestPlayer_Evolve1(t *testing.T) {
	inputs := []UserInputEventMessage{
		{Timestamp: 0, W: false, A: false, S: false, D: false},
		{Timestamp: 100, W: true, A: false, S: false, D: false},
		{Timestamp: 200, W: false, A: true, S: false, D: false},
		{Timestamp: 300, W: false, A: false, S: true, D: false},
		{Timestamp: 400, W: false, A: false, S: false, D: true},
	}

	InputStack := InputStack{Inputs: inputs}

	player := Player{
		InputStack: InputStack,
		Pos:        [2]float64{0, 0},
	}

	player.Evolve(500)

	if player.Pos[0] != 0 || player.Pos[1] != 0 {
		t.Errorf("Expected player position to be (0, 0), got (%f, %f)", player.Pos[0], player.Pos[1])
	}
}
