package hub

type Game struct {
	Player1 *Client
	Player2 *Client
}

type Hub struct {
	// Games that are registered
	Games 			 map[Game]GameState
	register		 <-chan *GameState
	broadcast 	     chan<- *GameState
	receive          <-chan *GameState
}

func NewHub() *Hub {
	return &Hub{
		Games:      make(map[Game]GameState),
		register:   make(<-chan *GameState),
		broadcast:  make(chan<- *GameState),
		receive:    make(<-chan *GameState),
	}
}

func (h *Hub) Run() {
	for {
		break
	}
}