package hub

import (
	"fmt"
	"log"
)

type User struct {
	Username       string
	HashedPassword string
}

// Hub stores all the ongoing games and deals with registrations and validation
// of player moves.
type Hub struct {
	// Games that are registered

	Games map[*Game]*GameState

	Register   chan IPlayer
	Unregister chan IPlayer
	MakeTurn   chan Turn

	Users []User
}

// NewHub sets up a Hub and returns the memory location.
func NewHub() *Hub {
	return &Hub{
		Games: make(map[*Game]*GameState),

		Register:   make(chan IPlayer),
		Unregister: make(chan IPlayer),
		MakeTurn:   make(chan Turn),
	}
}

// AddToGameOrNewGame either adds a client to the first available game or creates a
// new game and adds them to that.
func (h *Hub) AddToGameOrNewGame(player IPlayer) error {

	switch p := player.(type) {
	case *Player:
		// Add to the relevant map of games.
		for game := range h.Games {
			if game.SlotsFree() > 0 {
				err := game.AddClient(p)
				//p.Stream <- h.Games[p.Game].GameStateOutput()
				if err != nil {
					return fmt.Errorf("err when adding player to hub: %s", err.Error())
				}
				player.(*Player).Pos = [2]float64{100., 100.}
				return nil
			}
		}

		newGame := &Game{Players: []*Player{p}, Status: GameWaiting, t0: 0.}
		h.Games[newGame] = &GameState{}
		player.(*Player).Game = newGame
		//player.(*Player).Stream <- h.Games[player.(*Player).Game].GameStateOutput()
		return nil

	default:
		return fmt.Errorf("AddToGameOrNewGame type error")
	}

}

func remove(s []*Player, i int) []*Player {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

func key(s []*Player, player *Player) (int, error) {
	for i, plyr := range s {
		if plyr == player {
			return i, nil
		}
	}

	return -1, fmt.Errorf("player does not exist")
}

func (h *Hub) UnregisterClient(player IPlayer) error {
	// Search for this client and remove it.

	var unregErr error

	switch p := player.(type) {
	case *Player:
		for game := range h.Games {

			playerKey, keyErr := key(game.Players, p)
			if keyErr != nil {
				unregErr = keyErr
				continue
			}

			remove(game.Players, playerKey)
			return nil
		}
	default:
		fmt.Errorf("UnregisterClient type error")
	}
	return unregErr
}

func (h *Hub) Run() {
	for {
		select {
		case x, ok := <-h.Register:
			// Register incoming new client, add them to a lobby
			if ok {
				err := h.AddToGameOrNewGame(x)
				if err != nil {
					break
				}
			} else {
				return
			}

		case x, ok := <-h.Unregister:
			// Unregister and clean up
			if ok {
				err := h.UnregisterClient(x)
				if err != nil {
					break
				}
			} else {
				return
			}

		case turn, ok := <-h.MakeTurn:
			// Process a player turn
			if ok {
				switch p := turn.player.(type) {
				case *Player:
					turnErr := p.ProcessTurn(turn.encodedTurn)
					if turnErr != nil {
						log.Printf("was unable to process turn for player %q\n", p)
					}
				}
			} else {
				return
			}
		}
	}
}
