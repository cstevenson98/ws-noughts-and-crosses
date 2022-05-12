package hub

import (
	"time"
)

func (h *Hub) PlayerCount() (sum int) {
	for game := range h.Games {
		sum += len(game.Players)
	}
	return
}

func (h *Hub) LogOnInterval(interval time.Duration) {
}