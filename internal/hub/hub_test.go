package hub

import (
	"testing"
)

func TestGame_SlotsFree(t *testing.T) {
	mockGame2 := Game{Player1: nil, Player2: nil}
	mockGame1 := Game{Player1: &Client{}, Player2: nil}
	mockGame0 := Game{Player1: &Client{}, Player2: &Client{}}

	if mockGame0.SlotsFree() != 0 {
		t.Errorf("test failed: must be 0")
	} else if mockGame1.SlotsFree() != 1 {
		t.Errorf("test failed: must be 1")
	} else if mockGame2.SlotsFree() != 2 {
		t.Errorf("test failed: must be 2")
	}
}

func TestHub_AddToGameOrNewGame(t *testing.T) {
	// Completely empty hub. New game with client added in
	mockHub := NewHub()
	err := mockHub.AddToGameOrNewGame(&Client{mockHub, nil, nil})
	if err != nil {
		t.Errorf("Test failed unexpectedly: error adding client")
	}
	// A hub which consists of one player already waiting in game: add player to the free slot
	mockHub2 := &Hub{
		Games: map[*Game]GameState{
			&Game{&Client{}, nil} : GameState{},
		},
	}
	err = mockHub2.AddToGameOrNewGame(&Client{mockHub, nil, nil})
	if err != nil {
		t.Errorf("Test failed unexpectedly: error adding client")
	}
	// A hub which already has one game which is full, so create new game and add client to that
	mockHub3 := &Hub{
		Games: map[*Game]GameState{
			&Game{&Client{}, &Client{}} : GameState{},
		},
	}
	err = mockHub3.AddToGameOrNewGame(&Client{mockHub, nil, nil})
	if err != nil {
		t.Errorf("Test failed unexpectedly: error adding client")
	}
}