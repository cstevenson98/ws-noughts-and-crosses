package hub

import (
	"encoding/json"
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
	for range time.Tick(interval) {
		for game := range h.Games {
			for _, player := range game.Players {
				payload, _ := json.Marshal(player.Pos)
				player.Stream <- payload
			}
			game.t0 += interval.Seconds()
		}
	}
}
