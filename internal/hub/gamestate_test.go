package hub

import (
	"fmt"
	"testing"
)

func TestGameState_BoardToOutput(t *testing.T) {
	state := *NewState()
	fmt.Println(string(state.BoardToOutput()))
}