package hub

import (
	"fmt"
	"log"
)

// Hub stores all the ongoing games and deals with registrations and validation
// of player moves.
type Hub struct {
	// Games that are registered

	OandXGames map[*OandXGame]*OandXState

	Register   chan IPlayer
	Unregister chan IPlayer
	MakeTurn   chan Turn
}

// NewHub sets up a Hub and returns the memory location.
func NewHub() *Hub {
	return &Hub{
		OandXGames: make(map[*OandXGame]*OandXState),

		Register:   make(chan IPlayer),
		Unregister: make(chan IPlayer),
		MakeTurn:   make(chan Turn),
	}
}

// AddToGameOrNewGame either adds a client to the first available game or creates a
// new game and adds them to that.
func (h *Hub) AddToGameOrNewGame(player IPlayer) error {

	switch p := player.(type) {
	case *OandXPlayer:
		// Add to the relevant map of games.
		for game := range h.OandXGames {
			if game.SlotsFree() > 0 {
				err := game.AddClient(player.(*OandXPlayer))
				p.Stream <- h.OandXGames[p.Game].BoardToOutput()
				if err != nil {
					return fmt.Errorf("err when adding player to hub: %s", err.Error())
				}
				return nil
			}
		}

		newGame := &OandXGame{Player1: player.(*OandXPlayer), Player2: nil, Status: GameWaiting}
		h.OandXGames[newGame] = &OandXState{}
		h.OandXGames[newGame].Clear()
		player.(*OandXPlayer).Game = newGame
		player.(*OandXPlayer).Stream <- h.OandXGames[player.(*OandXPlayer).Game].BoardToOutput()
		return nil

	default:
		return fmt.Errorf("AddToGameOrNewGame type error")
	}

}

func (h *Hub) UnregisterClient(player IPlayer) error {
	// Search for this client and remove it.
	switch p := player.(type) {
	case *OandXPlayer:
		for game := range h.OandXGames {
			if game.Player1 == p {
				game.Player1 = nil
				if game.Player1 == nil && game.Player2 == nil {
					delete(h.OandXGames, game)
				}
				return nil
			} else if game.Player2 == player {
				game.Player2 = nil
				if game.Player1 == nil && game.Player2 == nil {
					delete(h.OandXGames, game)
				}
				return nil
			}
		}
	default:
		fmt.Errorf("UnregisterClient type error")
	}
	return fmt.Errorf("player was not found registered in hub")
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
				case *OandXPlayer:
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