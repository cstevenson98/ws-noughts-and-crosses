package hub

import (
	"time"
)

func (h *Hub) PlayerCount() (sum int) {
	//for game, _ := range h.Games {
	//	if game.Player1 != nil {
	//		sum += 1
	//	}
	//	if game.Player2 != nil {
	//		sum += 1
	//	}
	//}
	return
}

func (h *Hub) LogOnInterval(interval time.Duration) {
	//for range time.Tick(interval) {
	//	log.Printf("There are currently %v active game(s), with %v total player(s)", len(h.Games), h.PlayerCount())
	//}
}