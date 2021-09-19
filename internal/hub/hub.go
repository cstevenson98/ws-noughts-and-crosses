package hub

import "fmt"

type Game struct {
	Player1 *Client
	Player2 *Client
}

// AddClient adds a client to a game in the first available slot
func (g *Game) AddClient(client *Client) error {
	if g.Player1 == nil {
		g.Player1 = client
	} else if g.Player2 == nil {
		g.Player2 = client
	} else {
		return fmt.Errorf("no slots in game")
	}
	return nil
}

// SlotsFree returns the number of unoccupied slots in a game.
func (g *Game) SlotsFree() (slots int) {
	if g.Player1 == nil {
		slots += 1
	}
	if g.Player2 == nil {
		slots += 1
	}
	return
}

type Turn struct {
	client *Client
	move   [2]int
}

// Hub stores all the ongoing games and deals with registrations and validation of player moves.
type Hub struct {
	// Games that are registered
	Games 			 map[*Game]GameState
	Register		 chan *Client
	Unregister		 chan *Client
	MakeTurn		 chan Turn
}

// NewHub sets up a Hub and returns the memory location.
func NewHub() *Hub {
	return &Hub{
		Games:      make(map[*Game]GameState),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		MakeTurn:   make(chan Turn),
	}
}

// AddToGameOrNewGame either adds a client to the first available game or creates a
// new game and adds them to that.
func (h *Hub) AddToGameOrNewGame(client *Client) error {
	for game, _ := range h.Games {
		if game.SlotsFree() > 0 {
			err := game.AddClient(client)
			if err != nil {
				return fmt.Errorf("err when adding client to hub: %s", err.Error())
			} else {
				return nil
			}
		}
	}

	newGame := &Game{client, nil}
	h.Games[newGame] = GameState{}
	return nil
}

func (h *Hub) UnregisterClient(client *Client) error {
	// Search for this client and remove it.
	for game, _ := range h.Games {
		if game.Player1 == client {
			game.Player1 = nil
		} else if game.Player2 == client {
			game.Player2 = nil
		}
		if game.Player1 == nil && game.Player2 == nil {
			delete(h.Games, game)
		}
		return nil
	}
	return fmt.Errorf("client was not found registered in hub")
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
			// Register incoming new client, add them to a lobby
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
				fmt.Printf("%v\n", *turn.client)
			} else {
				return
			}
		}
	}
}