package hub

import (
	"fmt"
	"log"
)

type User struct {
	Username       string
	HashedPassword string
}

type Hub struct {
	Games map[*Game]bool

	Register   chan IPlayer
	Unregister chan IPlayer
	MakeTurn   chan Turn

	Users []User
}

func NewHub() *Hub {
	return &Hub{
		Games: make(map[*Game]bool),

		Register:   make(chan IPlayer),
		Unregister: make(chan IPlayer),
		MakeTurn:   make(chan Turn),
	}
}

func (h *Hub) AddToGameOrNewGame(player IPlayer) error {

	switch p := player.(type) {
	case *Player:
		// Add to the relevant map of games.
		for game := range h.Games {
			if game.SlotsFree() > 0 {
				err := game.AddClient(p)
				if err != nil {
					return fmt.Errorf("err when adding player to hub: %s", err.Error())
				}
				return nil
			}
		}

		playerMap := make(map[*Player]bool)
		playerMap[p] = true

		newGame := &Game{Players: playerMap, Status: GameWaiting, t0: 0., dt: dtDefault}
		player.(*Player).Game = newGame

		go newGame.RunGame()
		
		return nil

	default:
		return fmt.Errorf("AddToGameOrNewGame type error")
	}

}

func (h *Hub) UnregisterClient(player IPlayer) error {
	switch p := player.(type) {
	case *Player:
		for game := range h.Games {
			if _, exists := game.Players[p]; exists {
				delete(game.Players, p)
			}
		}
	default:
		fmt.Errorf("UnregisterClient type error")
	}
	return fmt.Errorf("player not found")
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
						log.Printf("was unable to process turn for player %v\n", p)
					}
				}
			} else {
				return
			}
		}
	}
}