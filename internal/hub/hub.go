package hub

import (
	"fmt"
)

// Hub stores all the ongoing games and deals with registrations and validation
// of player moves.
type Hub struct {
	// Games that are registered
	Games      map[*Game]*GameState
	Register   chan *Player
	Unregister chan *Player
	MakeTurn   chan Turn
}

// NewHub sets up a Hub and returns the memory location.
func NewHub() *Hub {
	return &Hub{
		Games:      make(map[*Game]*GameState),
		Register:   make(chan *Player),
		Unregister: make(chan *Player),
		MakeTurn:   make(chan Turn),
	}
}

// AddToGameOrNewGame either adds a client to the first available game or creates a
// new game and adds them to that.
func (h *Hub) AddToGameOrNewGame(player *Player) error {
	for game := range h.Games {
		if game.SlotsFree() > 0 {
			err := game.AddClient(player)
			player.Stream <- h.Games[player.Game].BoardToOutput()
			if err != nil {
				return fmt.Errorf("err when adding player to hub: %s", err.Error())
			}
			return nil
		}
	}

	newGame := &Game{player, nil, GameWaiting}
	h.Games[newGame] = &GameState{}
	h.Games[newGame].Clear()
	player.Game = newGame
	player.Stream <- h.Games[player.Game].BoardToOutput()
	return nil
}

func (h *Hub) UnregisterClient(player *Player) error {
	// Search for this client and remove it.
	for game := range h.Games {
		if game.Player1 == player {
			game.Player1 = nil
			if game.Player1 == nil && game.Player2 == nil {
				delete(h.Games, game)
			}
			return nil
		} else if game.Player2 == player {
			game.Player2 = nil
			if game.Player1 == nil && game.Player2 == nil {
				delete(h.Games, game)
			}
			return nil
		}
	}
	return fmt.Errorf("player was not found registered in hub")
}

// ProcessTurn takes the turn information pumped up from the player
// read routine and determines what can be done with, and enacts the
// appropriate action.
func (h *Hub) ProcessTurn(turn Turn) error {

	var nextBoard = h.Games[turn.player.Game]
	playerLabel := turn.player.Game.WhichPlayer(turn.player)
	var otherPlayer *Player
	if playerLabel == GamePlayer1 {
		nextBoard.Board[turn.move[0]][turn.move[1]] = "X"
		otherPlayer = turn.player.Game.Player2
	} else if playerLabel == GamePlayer2 {
		nextBoard.Board[turn.move[0]][turn.move[1]] = "0"
		otherPlayer = turn.player.Game.Player1
	} else {
		return fmt.Errorf("error: unknown player cannot make turn")
	}

	output := nextBoard.BoardToOutput()
	turn.player.Stream <- output
	if otherPlayer != nil {
		otherPlayer.Stream <- output
	}
	h.Games[turn.player.Game] = nextBoard

	return nil
}

// Run is the function which deals with the core features of the websocket hub.
// It registers new clients and adds them to a game. It also deals with the players'
// turns.
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
				h.ProcessTurn(turn)
			} else {
				return
			}
		}
	}
}
